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

	// OK
	params := objectia.PDFCreateParams{
		DocumentHTML: "<html>This is a test, too</html>",
	}
	pdfDoc, err := client.PDF.Create(&params)
	assert.NoError(t, err)
	assert.NotEmpty(t, pdfDoc)

	ioutil.WriteFile("/tmp/gogo.pdf", pdfDoc, 0644)
}
