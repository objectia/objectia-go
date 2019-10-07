package objectia

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// Parameters struct
type Parameters struct {
	contentType string
	params      map[string]interface{}
	files       map[string]string
}

// NewParameters creates a new parameters object
func NewParameters() *Parameters {
	return &Parameters{
		params: make(map[string]interface{}),
		files:  make(map[string]string),
	}
}

// Add adds a new key/value pair to payload.
func (p *Parameters) Add(key string, value interface{}) {
	p.params[key] = value
}

// AddFile adds a new file to be uploaded
func (p *Parameters) AddFile(key, file string) {
	p.files[key] = file
}

// GetContentType returns the content type of the payload
func (p *Parameters) GetContentType() string {
	return p.contentType
}

// Encode marshalls the payload into JSON.
func (p *Parameters) Encode() (*bytes.Buffer, error) {
	result := &bytes.Buffer{}

	if len(p.files) > 0 {
		// Has file attachments

		writer := multipart.NewWriter(result)
		defer writer.Close()

		// The files...
		for key, fn := range p.files {
			file, err := os.Open(fn)
			if err != nil {
				return nil, err
			}
			defer file.Close()
			part, err := writer.CreateFormFile(key, filepath.Base(fn))
			if err != nil {
				return nil, err
			}
			io.Copy(part, file)
		}

		// The other attributes
		for key, val := range p.params {
			part, err := writer.CreateFormField(key)
			if err != nil {
				return nil, err
			}

			b, _ := getBytes(val)
			part.Write(b)

			fmt.Println(b)
		}

		p.contentType = writer.FormDataContentType()
		fmt.Println("CONTENT TYPE: ", p.contentType)
		writer.Close()
	} else {
		// No files, only key/value attributes
		p.contentType = "application/json"
		err := json.NewEncoder(result).Encode(p.params)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
