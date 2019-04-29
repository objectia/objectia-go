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
	assert.Equal(t, 12345, result.GeoLocation)
}

/*
func Test_GeoLocation_Get(t *testing.T) {
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	result, err := client.GeoLocation.Get("8.8.8.8")
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "US", result.CountryCode)
}

func Test_GeoLocation_GetBulk(t *testing.T) {
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	results, err := client.GeoLocation.GetBulk([]string{"8.8.8.8", "google.com"})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(results))
	for _, v := range results {
		assert.NotNil(t, v)
		assert.Equal(t, "US", v.CountryCode)
	}
}*/
