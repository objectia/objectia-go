package objectia

import (
	"bytes"
	"encoding/json"
	"io"
)

// Parameters struct
type Parameters struct {
	params map[string]interface{}
}

// NewParameters creates a new parameters object
func NewParameters() *Parameters {
	return &Parameters{
		params: make(map[string]interface{}),
	}
}

// Add adds a new key/value pair to payload.
func (p *Parameters) Add(key string, value interface{}) {
	p.params[key] = value
}

// Encode marshalls the payload into JSON.
func (p *Parameters) Encode() (rw io.ReadWriter, err error) {
	rw = new(bytes.Buffer)
	err = json.NewEncoder(rw).Encode(p.params)
	return
}
