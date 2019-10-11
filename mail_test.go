package objectia_test

import (
	"testing"

	"github.com/objectia/objectia-go"
	"github.com/stretchr/testify/assert"
)

func Test_Mail_Send(t *testing.T) {
	apiKey := "d562d9b7c24b47f8a1cd8689db0404fc"
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	// OK
	m := objectia.NewMessage("ok@demo2.org", "Test", "This is just a test", "ok@demo2.org")
	messageID, err := client.Mail.Send(m)
	assert.NoError(t, err)
	assert.NotEmpty(t, messageID)
}

func Test_Mail_Send_Hard_Bounce(t *testing.T) {
	apiKey := "d562d9b7c24b47f8a1cd8689db0404fc"
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	// Hard bounce - permanent error
	m := objectia.NewMessage("bounce@demo2.org", "Test", "This is just a test", "ok@demo2.org")
	messageID, err := client.Mail.Send(m)
	assert.NoError(t, err)
	assert.NotEmpty(t, messageID)

}

func Test_Mail_Send_Graylisted(t *testing.T) {
	apiKey := "d562d9b7c24b47f8a1cd8689db0404fc"
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	// Sender gray listed - transient error
	m := objectia.NewMessage("gray@demo2.org", "Test", "This is just a test", "ok@demo2.org")
	messageID, err := client.Mail.Send(m)
	assert.NoError(t, err)
	assert.NotEmpty(t, messageID)

}

func Test_Mail_Send_Mailbox_Full(t *testing.T) {
	apiKey := "d562d9b7c24b47f8a1cd8689db0404fc"
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	// Mailbox temporary full
	m := objectia.NewMessage("ok@demo2.org", "Test", "This is just a test", "full@demo2.org")
	messageID, err := client.Mail.Send(m)
	assert.NoError(t, err)
	assert.NotEmpty(t, messageID)

}

func Test_Mail_Send_Unknown_Recipient(t *testing.T) {
	apiKey := "d562d9b7c24b47f8a1cd8689db0404fc"
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	// Mailbox not found - permanment error
	m := objectia.NewMessage("ok@demo2.org", "Test", "This is just a test", "none@demo2.org")
	messageID, err := client.Mail.Send(m)
	assert.NoError(t, err)
	assert.NotEmpty(t, messageID)

}

func Test_Mail_Send_DomainError(t *testing.T) {
	apiKey := "d562d9b7c24b47f8a1cd8689db0404fc"
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	// Domain not found
	m := objectia.NewMessage("ok@demo2.org", "Test", "This is just a test", "none@demo2.xxx")
	messageID, err := client.Mail.Send(m)
	assert.NoError(t, err)
	assert.NotEmpty(t, messageID)
}
