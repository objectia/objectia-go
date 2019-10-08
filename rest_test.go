package objectia

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var apiKey string

func init() {
	apiKey = os.Getenv("OBJECTIA_APIKEY")
	if len(apiKey) == 0 {
		log.Fatalln("OBJECTIA_APIKEY environment variable not defined")
	}
}

func Test_Get(t *testing.T) {
	client, err := NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	client.RetryWaitMax = 10 * time.Second
	client.Logger = log.New(os.Stderr, "", log.LstdFlags)

	var resp Response
	_, err = client.get("/v1/test", nil, &resp)
	assert.NoError(t, err)
}

func Test_Post(t *testing.T) {
	client, err := NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	params := NewParameters()
	params.Add("name", "Otto")

	var resp Response
	err = client.post("/v1/test", params, &resp)
	assert.NoError(t, err)
}

func Test_Post_File(t *testing.T) {
	client, err := NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	params := NewParameters()
	params.AddFile("me.png", "/users/Otto/me.png")
	params.AddFile("me2.png", "/users/Otto/me2.png")
	params.Add("name", "dsadsdasd")

	var resp Response
	err = client.post("/v1/test", params, &resp)
	assert.NoError(t, err)
}
