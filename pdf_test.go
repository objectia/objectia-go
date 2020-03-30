package objectia_test

import (
	"io/ioutil"
	"testing"

	"github.com/objectia/objectia-go"
	"github.com/stretchr/testify/assert"
)

func Test_PDF_HTML(t *testing.T) {
	client, err := objectia.NewClient(apiKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	params := objectia.PDFCreateParams{
		DocumentHTML: "<html><h1>Hello world</h1>This is a test of the <strong>PDF API</strong></html>",
	}
	buf, err := client.PDF.Create(&params)
	assert.NoError(t, err)
	assert.NotEmpty(t, buf)

	ioutil.WriteFile("/tmp/pdftest.pdf", buf, 0644)
}
