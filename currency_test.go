package objectia_test

import (
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/objectia/objectia-go"
	"github.com/stretchr/testify/assert"
)

func Test_Currency_Latest(t *testing.T) {
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	res, err := client.Currency.GetLatest(nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	//***************************************************
	// with options
	options := objectia.CurrencyRateOptions{
		Base:       "EUR",
		Currencies: []string{"EUR", "GBP", "USD"},
	}
	res, err = client.Currency.GetLatest(&options)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func Test_Currency_Historical(t *testing.T) {
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)

	res, err := client.Currency.GetHistorical(yesterday, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	//***************************************************
	// with options
	options := objectia.CurrencyRateOptions{
		Base:       "EUR",
		Currencies: []string{"EUR", "GBP", "USD"},
	}
	res, err = client.Currency.GetHistorical(yesterday, &options)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func Test_Currency_Convert(t *testing.T) {
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	now := time.Now()
	res, err := client.Currency.Convert(100.0, "USD", "EUR", now)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func Test_Currency_Change(t *testing.T) {
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)

	res, err := client.Currency.GetChange(yesterday, now, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func Test_Currency_Timeframe(t *testing.T) {
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)

	res, err := client.Currency.GetTimeframe("USD", yesterday, now, "EUR")
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func Test_Currency_List(t *testing.T) {
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	res, err := client.Currency.List()
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	spew.Dump(res.Currencies)
}
