package objectia

import (
	"net/http"
	"time"
)

const (
	version             = "1.0.0"
	userAgent           = "objectia-go/" + version
	apiBaseURL          = "https://api.objectia.com"
	defaultTimeout      = 30 * time.Second
	defaultRetryMax     = 4
	defaultRetryWaitMin = 1 * time.Second
	defaultRetryWaitMax = 30 * time.Second
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
	// Public APIs
	GeoLocation *GeoLocation
	Currency    *Currency
	PDF         *PDF
	Mail        *Mail
	SMS         *SMS
	Usage       *Usage
}

// GeoLocation api functions
type GeoLocation struct {
	client *Client
}

// Currency api functions
type Currency struct {
	client *Client
}

// PDF api functions
type PDF struct {
	client *Client
}

// Mail api functions
type Mail struct {
	client *Client
}

// SMS api functions
type SMS struct {
	client *Client
}

// Usage api functions
type Usage struct {
	client *Client
}

// NewClient creates a new Client with the provided apiKey and an optional httpClient.
func NewClient(apiKey string, httpClient *http.Client) (*Client, error) {
	if len(apiKey) == 0 {
		return nil, NewError("err-missing-apikey", "No API key provided")
	}

	c := &Client{
		apiBaseURL:   apiBaseURL,
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
	c.Currency = &Currency{client: c}
	c.PDF = &PDF{client: c}
	c.Mail = &Mail{client: c}
	c.SMS = &SMS{client: c}
	c.Usage = &Usage{client: c}

	return c, nil
}

// GetVersion returns the client version string.
func (c *Client) GetVersion() string {
	return version
}
