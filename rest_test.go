package objectia

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Get(t *testing.T) {
	client, err := NewClient("test", nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	client.RetryWaitMax = 10 * time.Second

	client.Logger = log.New(os.Stderr, "", log.LstdFlags)

	var resp Response
	_, err = client.get("/v1/geoip/8.8.8.8", nil, &resp)
	assert.Error(t, err)
}
