package memory

import (
	"encoding/json"

	"github.com/anothermemory/storage"
	"github.com/anothermemory/storagedir"
	"github.com/pkg/errors"
)

type storageTypeJSON struct {
	Type string `json:"type"`
}

// UnmarshalStorageJSONFunc represents a function which is able to unmarshal storage from json representation
type UnmarshalStorageJSONFunc func(b []byte) (storage.Interface, error)

// UnmarshalStorageJSON is default implementation of UnmarshalStorageJSONFunc
func UnmarshalStorageJSON(b []byte) (storage.Interface, error) {
	var j storageTypeJSON
	err := json.Unmarshal(b, &j)

	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal storage type")
	}

	var s storage.Interface
	if j.Type == storagedir.TypeDirectory || j.Type == storagedir.TypeDirectoryInMemory {
		s, err = storagedir.NewDirectoryStorageFromJSONConfig(b)
		if nil != err {
			return nil, errors.Wrap(err, "Failed to unmarshal storage")
		}
	}

	if nil == s {
		return nil, errors.Errorf("Cannot load storage. Unsupported storage type received: %s", j.Type)
	}

	return s, nil
}
