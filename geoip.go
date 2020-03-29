package objectia

import (
	"fmt"
	"strings"

	"github.com/objectia/objectia-go/internal"
)

// Get retrieves the geolocation for the given domain or IP address
func (c *GeoLocation) Get(ip string, options *GeoLocationOptions) (*IPLocation, error) {
	if len(ip) == 0 {
		return nil, NewError("err-invalid-ip", "Invalid IP address")
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
	err = internal.FromMap(resp.Data, result)
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
		return nil, NewError("err-invalid-ip", "Invalid IP address")
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
	err = internal.FromMap(resp.Data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
