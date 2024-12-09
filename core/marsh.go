package core

import (
	"encoding/json"
)

func (c *JSONCore) UnmarshalJSON(b []byte) error {
	var err error = nil
	// unmarshal to array or obj
	arr := make([]interface{}, 1)
	obj := make(map[string]interface{})
	err = json.Unmarshal(b, &arr)

	if err != nil {
		// failed to parse into array
		err = json.Unmarshal(b, &obj)
		if err != nil {
			// failed to parse into array and object
			return err
		}
	}
	if arr[0] != nil {
		c.subject = arr
	}
	if len(obj) != 0 {
		c.subject = obj
	}
	return err
}

func (c *JSONCore) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.subject)
}
