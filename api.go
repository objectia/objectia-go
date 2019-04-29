package objectia

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	version        = "0.9.2"
	userAgent      = "objectia-go/" + version
	apiBaseURL     = "https://api.objectia.com/rest"
	defaultTimeout = time.Second * 30
)

// Connection errors
var (
	ErrConnectionTimedout = errors.New("Connection timed out")
	ErrConnectionRefused  = errors.New("Connection refused")
	ErrUknownHost         = errors.New("Unknown host")
	ErrNotModified        = errors.New("Not Modified")
	ErrInvalidIPAddress   = errors.New("Invalid IP address")
)

// Client encapsulates the api functions - must be created with NewClient()
type Client struct {
	apiKey     string
	apiBaseURL string
	httpClient *http.Client
	// Public APIs:
	GeoLocation *GeoLocation
	Usage       *Usage
}

// GeoLocation api functions
type GeoLocation struct {
	client *Client
}

// Usage api functions
type Usage struct {
	client *Client
}

// NewClient creates a new Client with the provided apiKey and an optional httpClient.
func NewClient(apiKey string, httpClient *http.Client) (*Client, error) {
	if len(apiKey) == 0 {
		return nil, errors.New("No API key provided")
	}

	c := &Client{
		apiBaseURL: apiBaseURL,
		apiKey:     apiKey,
		httpClient: httpClient,
	}

	// Use the default http client
	if c.httpClient == nil {
		c.httpClient = &http.Client{
			Timeout: defaultTimeout,
		}
	}

	// Attach the APIs
	c.GeoLocation = &GeoLocation{client: c}
	c.Usage = &Usage{client: c}

	return c, nil
}

// GetVersion returns the client version string.
func (c *Client) GetVersion() string {
	return version
}

// Get retrieves the geolocation for the given domain or IP address
func (c *GeoLocation) Get(ip string) (*IPLocation, error) {
	if len(ip) == 0 {
		return nil, ErrInvalidIPAddress
	}

	var resp Response
	_, err := c.client.get(fmt.Sprintf("/v1/geoip/%s", ip), nil, &resp)
	if err != nil {
		return nil, err
	}

	result := &IPLocation{}
	err = fromMap(resp.Data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetCurrent retrieves the geolocation for requester.
func (c *GeoLocation) GetCurrent() (*IPLocation, error) {
	var resp Response
	_, err := c.client.get("/v1/geoip/me", nil, &resp)
	if err != nil {
		return nil, err
	}

	result := &IPLocation{}
	err = fromMap(resp.Data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetBulk retrieves the geolocation for multiple domain names or IP addresses.
func (c *GeoLocation) GetBulk(iplist []string) ([]IPLocation, error) {
	var resp Response

	if len(iplist) == 0 {
		return nil, ErrInvalidIPAddress
	}
	ips := strings.Join(iplist, ",")

	_, err := c.client.get(fmt.Sprintf("/v1/geoip/%s", ips), nil, &resp)
	if err != nil {
		return nil, err
	}

	result := []IPLocation{}
	err = fromMap(resp.Data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Get returns the API usage for current month.
func (c *Usage) Get() (*APIUsage, error) {
	var resp Response
	_, err := c.client.get("/v1/usage", nil, &resp)
	if err != nil {
		return nil, err
	}

	result := &APIUsage{}
	err = fromMap(resp.Data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
