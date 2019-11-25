package objectia

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net"
	"net/http"
	"net/url"
	"time"
)

const timestampFormat = "Mon, Jan 2 2006 15:04:05 MST"

// Response model
type Response struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Code    string      `json:"code"`
	Data    interface{} `json:"data"`
}

func (c *Client) get(path string, oldEtag *ETag, result interface{}) (*ETag, error) {
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	// Append old ETag to headers
	if oldEtag != nil {
		req.Header.Set("If-None-Match", oldEtag.Tag)
		req.Header.Set("If-Modified-Since", oldEtag.LastModified.Format(timestampFormat))
	}

	resp, err := c.execute(req, result)
	//fmt.Println("StatusCode:", resp.StatusCode)

	etag := new(ETag)
	if resp != nil {
		etag.Tag = resp.Header.Get("ETag")
		t, err := time.Parse(timestampFormat, resp.Header.Get("Last-Modified"))
		if err == nil {
			etag.LastModified = t
		}
	}
	return etag, err
}

func (c *Client) post(path string, params *Parameters, result interface{}) error {
	req, err := c.newRequest("POST", path, params)
	if err != nil {
		return err
	}
	_, err = c.execute(req, result)
	return err
}

func (c *Client) put(path string, params *Parameters, result interface{}) error {
	req, err := c.newRequest("PUT", path, params)
	if err != nil {
		return err
	}
	_, err = c.execute(req, result)
	return err
}

func (c *Client) delete(path string, result interface{}) error {
	req, err := c.newRequest("DELETE", path, nil)
	if err != nil {
		return err
	}
	_, err = c.execute(req, result)
	return err
}

func (c *Client) newRequest(method, path string, params *Parameters) (*http.Request, error) {
	var body io.ReadWriter
	body = nil
	var err error
	if params != nil {
		body, err = params.Encode()
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, c.apiBaseURL+path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", userAgent)
	if params != nil {
		req.Header.Set("Content-Type", params.GetContentType())
	}
	return req, nil
}

func (c *Client) execute(req *http.Request, result interface{}) (*http.Response, error) {
	if c.Logger != nil {
		c.Logger.Printf("[DEBUG] %s %s", req.Method, req.URL)
	}

	var resp *http.Response
	var err error

	for i := 0; ; i++ {
		var statusCode int

		// Attempt the request
		resp, err = c.httpClient.Do(req)
		if resp != nil {
			statusCode = resp.StatusCode
		}

		// Check if we have to retry
		retry, checkErr := checkRetry(req.Context(), resp, err)

		if err != nil {
			if c.Logger != nil {
				c.Logger.Printf("[ERROR] %s %s request failed: %v", req.Method, req.URL, err)
			}
		}

		if !retry {
			if checkErr != nil {
				err = checkErr
			} else {
				if statusCode == 200 || statusCode == 201 {
					err = parseResponse(resp, result)
				} else {
					err = parseErrorResponse(resp)
				}
			}
			return resp, err
		}

		// Make another attmpt?
		attemptsLeft := c.RetryMax - i
		if attemptsLeft <= 0 {
			break
		}

		wait := backoff(c.RetryWaitMin, c.RetryWaitMax, i)
		desc := fmt.Sprintf("%s %s", req.Method, req.URL)
		if statusCode > 0 {
			desc = fmt.Sprintf("%s (status: %d)", desc, statusCode)
		}
		if c.Logger != nil {
			c.Logger.Printf("[DEBUG] %s: retrying in %s (%d left)", desc, wait, attemptsLeft)
		}
		select {
		case <-req.Context().Done():
			return nil, req.Context().Err()
		case <-time.After(wait):
		}
	}

	if resp != nil {
		resp.Body.Close()
	}

	return nil, fmt.Errorf("%s %s giving up after %d attempts", req.Method, req.URL, c.RetryMax+1)
}

func checkConnectError(err error) error {
	switch t := err.(type) {
	case *url.Error:
		if err, ok := t.Err.(net.Error); ok && err.Timeout() {
			return NewError("err-timeout", "Connection timed out")
		}
		if _, ok := t.Err.(*net.OpError); ok {
			return NewError("err-connect", "Failed to connect")
		}
	}
	return err
}

func closeResponse(resp *http.Response) {
	err := resp.Body.Close()
	if err != nil {
	}
}

func parseResponse(resp *http.Response, result interface{}) error {
	if resp.Body == nil {
		return NewResponseError(resp.StatusCode, "err-no-data", "Server returned no data")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(body, result)
	if err != nil {
		err = NewResponseError(resp.StatusCode, "err-unexpected", "Unexpected response from server")
	}
	return err
}

func parseErrorResponse(resp *http.Response) error {
	if resp.Body == nil {
		return NewResponseError(resp.StatusCode, "err-no-data", "Server returned no data")
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var err error
	switch resp.StatusCode {
	case 304:
		err = NewResponseError(resp.StatusCode, "err-not-modified", "Not modified")
	case 500:
		err = NewResponseError(resp.StatusCode, "err-internal-server-error", "Internal server error")
	case 502:
		err = NewResponseError(resp.StatusCode, "err-bad-gateway", "Bad gateway")
	case 503:
		err = NewResponseError(resp.StatusCode, "err-service-unavailable", "Service unavailable")
	}
	if err != nil {
		return err
	}

	var result Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		err = NewResponseError(resp.StatusCode, "err-unexpected", "Unexpected response from server")
	} else if len(result.Message) == 0 {
		err = NewResponseError(resp.StatusCode, "err-unexpected", "Unexpected response from server")
	} else {
		err = NewResponseError(resp.StatusCode, result.Code, result.Message)
	}
	return err
}

// backoff will perform exponential backoff based on the attempt number and limited
// by the provided minimum and maximum durations.
func backoff(min, max time.Duration, attemptNum int) time.Duration {
	mult := math.Pow(2, float64(attemptNum)) * float64(min)
	sleep := time.Duration(mult)
	if float64(sleep) != mult || sleep > max {
		sleep = max
	}
	return sleep
}

// checkRetry checks if we should retry or not.
func checkRetry(ctx context.Context, resp *http.Response, err error) (bool, error) {
	// do not retry on context.Canceled or context.DeadlineExceeded
	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	if err != nil {
		return true, checkConnectError(err)
	}
	// Check the response code. We retry on 500-range responses to allow
	// the server time to recover, as 500's are typically not permanent
	// errors and may relate to outages on the server side. This will catch
	// invalid response codes as well, like 0 and 999.
	if resp.StatusCode == 0 || (resp.StatusCode >= 500 && resp.StatusCode != 501) {
		return true, nil
	}

	return false, nil
}
