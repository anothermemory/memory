package memory

import (
	"encoding/json"
)

// NewFromJSONConfig creates new memory instance from JSON configuration
func NewFromJSONConfig(b []byte) (Interface, error) {
	m := &Instance{unmarshalStorageJSONFunc: UnmarshalStorageJSON}

	err := json.Unmarshal(b, &m)

	if err != nil {
		return nil, err
	}

	return m, nil
}
