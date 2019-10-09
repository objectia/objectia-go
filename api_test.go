package objectia_test

import (
	"log"
	"os"
	"testing"

	"github.com/objectia/objectia-go"
	"github.com/stretchr/testify/assert"
)

var apiKey string

func init() {
	apiKey = os.Getenv("OBJECTIA_APIKEY")
	if len(apiKey) == 0 {
		log.Fatalln("OBJECTIA_APIKEY environment variable not defined")
	}
}

func Test_GetVersion(t *testing.T) {
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	v := client.GetVersion()
	assert.NotEmpty(t, v)
}

func Test_Usage_Get(t *testing.T) {
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	result, err := client.Usage.Get()
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 12345, result.GeoLocationRequests)
}

func Test_GeoLocation_Get(t *testing.T) {
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	result, err := client.GeoLocation.Get("8.8.8.8", nil)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "US", result.CountryCode)
}

func Test_GeoLocation_Get_with_options(t *testing.T) {
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	options := &objectia.GeoLocationOptions{
		DisplayFields: "country_code", // Return only country code
	}

	result, err := client.GeoLocation.Get("8.8.8.8", options)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "US", result.CountryCode)
}

func Test_GeoLocation_GetBulk(t *testing.T) {
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	results, err := client.GeoLocation.GetBulk([]string{"8.8.8.8", "google.com"}, nil)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(results))
	for _, v := range results {
		assert.NotNil(t, v)
		assert.Equal(t, "US", v.CountryCode)
	}
}

func Test_GeoLocation_GetCurrent(t *testing.T) {
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	result, err := client.GeoLocation.GetCurrent(nil)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func Test_GeoLocation_Get_Invalid_IP(t *testing.T) {
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	_, err = client.GeoLocation.Get("288.8.8.8", nil)
	assert.Error(t, err)
	if err != nil {
		e := err.(*objectia.Error)
		assert.Equal(t, "err-invalid-ip", e.Code)
	}
}

func Test_Mail_Send(t *testing.T) {
	apiKey := "d562d9b7c24b47f8a1cd8689db0404fc"
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	m := objectia.NewMessage("test@demo2.org", "Test", "This is just a test", "otto@doseth.com")
	messageID, err := client.Mail.Send(m)
	assert.NoError(t, err)
	assert.NotEmpty(t, messageID)
}
