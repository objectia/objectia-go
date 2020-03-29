package internal

import (
	"encoding/json"
)

// FromMap Converts from a map to a struct
func FromMap(data interface{}, result interface{}) error {
	m, _ := json.Marshal(data)
	return json.Unmarshal(m, &result)
}
