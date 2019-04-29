package objectia

import (
	"encoding/json"
)

// Convert from a map to a struct
func fromMap(data interface{}, result interface{}) error {
	m, _ := json.Marshal(data)
	return json.Unmarshal(m, &result)
}

/*func fromMap(data map[string]interface{}, result interface{}) error {
	m, _ := json.Marshal(data)
	return json.Unmarshal(m, &result)
}*/
