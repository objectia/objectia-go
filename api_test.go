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
