package objectia

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/objectia/objectia-go/internal"
)

// GetLatest retrieves the latest currency rates.
func (c *Currency) GetLatest(options *CurrencyRateOptions) (*CurrencyRate, error) {
	var resp Response

	// Process options
	query := ""
	if options != nil {
		query = "?base=" + url.QueryEscape(options.Base)
		if len(options.Currencies) > 0 {
			currencies := strings.Join(options.Currencies, ",")
			currencies = strings.Replace(currencies, " ", "", -1)
			query += "&currencies=" + currencies
		}
	}

	_, err := c.client.get(fmt.Sprintf("/v1/currency/latest%s", query), nil, &resp)
	if err != nil {
		return nil, err
	}

	result := &CurrencyRate{}
	err = internal.FromMap(resp.Data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetHistorical retrieves the currency rates for a given day in the past.
func (c *Currency) GetHistorical(date time.Time, options *CurrencyRateOptions) (*CurrencyRate, error) {
	var resp Response

	dt := date.Format("2006-01-02")

	// Process options
	query := ""
	if options != nil {
		query = "?base=" + url.QueryEscape(options.Base)
		if len(options.Currencies) > 0 {
			currencies := strings.Join(options.Currencies, ",")
			currencies = strings.Replace(currencies, " ", "", -1)
			query += "&currencies=" + currencies
		}
	}

	_, err := c.client.get(fmt.Sprintf("/v1/currency/historical/%s%s", dt, query), nil, &resp)
	if err != nil {
		return nil, err
	}

	result := &CurrencyRate{}
	err = internal.FromMap(resp.Data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Convert exchanges an amount from one currency to another on a given date.
func (c *Currency) Convert(amount float64, from, to string, date time.Time) (*CurrencyConvert, error) {
	var resp Response

	dt := date.Format("2006-01-02")

	_, err := c.client.get(fmt.Sprintf("/v1/currency/convert/%f/%s/%s?date=%s", amount, from, to, dt), nil, &resp)
	if err != nil {
		return nil, err
	}

	result := &CurrencyConvert{}
	err = internal.FromMap(resp.Data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetChange returns changes in currency rates between two dates.
func (c *Currency) GetChange(startDate, endDate time.Time, options *CurrencyRateOptions) (*CurrencyChange, error) {
	var resp Response

	start := startDate.Format("2006-01-02")
	end := endDate.Format("2006-01-02")

	// Process options
	query := ""
	if options != nil {
		query = "?base=" + url.QueryEscape(options.Base)
		if len(options.Currencies) > 0 {
			currencies := strings.Join(options.Currencies, ",")
			currencies = strings.Replace(currencies, " ", "", -1)
			query += "&currencies=" + currencies
		}
	}

	_, err := c.client.get(fmt.Sprintf("/v1/currency/change/%s/%s%s", start, end, query), nil, &resp)
	if err != nil {
		return nil, err
	}

	result := &CurrencyChange{}
	err = internal.FromMap(resp.Data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetTimeframe retrieves currency rate changes for each day in a give time period.
func (c *Currency) GetTimeframe(currency string, startDate, endDate time.Time, base string) (*CurrencyTimeframe, error) {
	var resp Response

	start := startDate.Format("2006-01-02")
	end := endDate.Format("2006-01-02")

	_, err := c.client.get(fmt.Sprintf("/v1/currency/timeframe/%s/%s/%s?base=%s", currency, start, end, base), nil, &resp)
	if err != nil {
		return nil, err
	}

	result := &CurrencyTimeframe{}
	err = internal.FromMap(resp.Data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// List retrieves a list of supported currencies.
func (c *Currency) List() (*CurrencyList, error) {
	var resp Response

	_, err := c.client.get("/v1/currency/list", nil, &resp)
	if err != nil {
		return nil, err
	}

	result := &CurrencyList{}
	err = internal.FromMap(resp.Data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
