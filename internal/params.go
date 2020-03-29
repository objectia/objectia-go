package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"reflect"
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

	if len(p.files) >= 0 {
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
			switch reflect.TypeOf(val).Kind() {
			case reflect.Slice, reflect.Array:
				v := reflect.ValueOf(val)
				for i := 0; i < v.Len(); i++ {
					s := fmt.Sprintf("%v", v.Index(i).Interface())
					writer.WriteField(key, s)
				}
			default:
				s := fmt.Sprintf("%v", val)
				writer.WriteField(key, s)
			}
		}

		p.contentType = writer.FormDataContentType()
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
