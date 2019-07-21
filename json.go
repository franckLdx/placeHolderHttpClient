package main

import (
	"encoding/json"
	"fmt"
)

func ParseToUnknownArray(data []byte) (*[]map[string]interface{}, error) {
	var parsed []map[string]interface{}
	if err := json.Unmarshal(data, &parsed); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal response")
	}
	return &parsed, nil
}
