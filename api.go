package objectia

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	version             = "1.0.0"
	userAgent           = "objectia-go/" + version
	apiBaseURL          = "https://api.objectia.com/rest"
	mockapiBaseURL      = "https://mock-api.objectia.com/rest"
	defaultTimeout      = 30 * time.Second
	defaultRetryMax     = 4
	defaultRetryWaitMin = 1 * time.Second
	defaultRetryWaitMax = 30 * time.Second
)

// Connection errors
var (
	ErrConnectionTimedout = errors.New("Connection timed out")
	ErrConnectionRefused  = errors.New("Connection refused")
	ErrUnknownHost        = errors.New("Unknown host")
	ErrNotModified        = errors.New("Not Modified")
	ErrInvalidIPAddress   = errors.New("Invalid IP address")
)

// Logger interface allows to use other loggers than standard log.Logger.
type Logger interface {
	Printf(string, ...interface{})
}

// Client encapsulates the api functions - must be created with NewClient()
type Client struct {
	apiKey     string
	apiBaseURL string
	httpClient *http.Client
	// Public properties
	Logger       Logger
	RetryMax     int
	RetryWaitMin time.Duration
	RetryWaitMax time.Duration
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

	baseURL := apiBaseURL
	if apiKey == "mock" {
		baseURL = mockapiBaseURL
	}

	c := &Client{
		apiBaseURL:   baseURL,
		apiKey:       apiKey,
		httpClient:   httpClient,
		RetryMax:     defaultRetryMax,
		RetryWaitMin: defaultRetryWaitMin,
		RetryWaitMax: defaultRetryWaitMax,
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

//************************************************

// Get retrieves the geolocation for the given domain or IP address
func (c *GeoLocation) Get(ip string, options *GeoLocationOptions) (*IPLocation, error) {
	if len(ip) == 0 {
		return nil, ErrInvalidIPAddress
	}

	// Process options
	query := ""
	if options != nil {
		query = "?fields=" + options.DisplayFields
		if options.LookupHostname {
			query += "&hostname=true"
		}
		if options.ShowSecurityInfo {
			query += "&security=true"
		}
	}

	var resp Response
	_, err := c.client.get(fmt.Sprintf("/v1/geoip/%s%s", ip, query), nil, &resp)
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

// GetCurrent retrieves the geolocation for the requester.
func (c *GeoLocation) GetCurrent(options *GeoLocationOptions) (*IPLocation, error) {
	return c.Get("myip", options)
}

// GetBulk retrieves the geolocation for multiple domain names or IP addresses.
func (c *GeoLocation) GetBulk(iplist []string, options *GeoLocationOptions) ([]IPLocation, error) {
	var resp Response

	if len(iplist) == 0 {
		return nil, ErrInvalidIPAddress
	}
	ips := strings.Join(iplist, ",")

	// Process options
	query := ""
	if options != nil {
		query = "?fields=" + options.DisplayFields
		if options.LookupHostname {
			query += "&hostname=true"
		}
		if options.ShowSecurityInfo {
			query += "&security=true"
		}
	}

	_, err := c.client.get(fmt.Sprintf("/v1/geoip/%s%s", ips, query), nil, &resp)
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

//************************************************

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
