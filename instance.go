package memory

import (
	"encoding/json"

	"github.com/anothermemory/storage"
	"github.com/anothermemory/unit"
	"github.com/pkg/errors"
)

// Instance represents memory instance
type Instance struct {
	name                     string
	root                     unit.Unit
	storage                  storage.Interface
	unmarshalStorageJSONFunc UnmarshalStorageJSONFunc
}

// New creates new memory instance
func New(s storage.Interface) Interface {
	return &Instance{storage: s, unmarshalStorageJSONFunc: UnmarshalStorageJSON}
}

// Name returns memory name
func (i *Instance) Name() string {
	return i.name
}

// SetName sets new memory name
func (i *Instance) SetName(n string) {
	i.name = n
}

// Root returns memory root unit
func (i *Instance) Root() unit.Unit {
	return i.root
}

// SetRoot sets new memory root unit
func (i *Instance) SetRoot(u unit.Unit) {
	i.root = u
}

// Storage returns memory storage
func (i *Instance) Storage() storage.Interface {
	return i.storage
}

// SetUnmarshalStorageJSONFunc sets new function which will perform storage unmarshal from config
func (i *Instance) SetUnmarshalStorageJSONFunc(f UnmarshalStorageJSONFunc) {
	i.unmarshalStorageJSONFunc = f
}

type instanceJSON struct {
	Name    string          `json:"name"`
	Root    string          `json:"root"`
	Storage json.RawMessage `json:"storage"`
}

// MarshalJSON converts memory instance to json representation
func (i *Instance) MarshalJSON() ([]byte, error) {
	bytes, err := json.Marshal(i.storage)

	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal memory storage")
	}
	return json.Marshal(instanceJSON{Name: i.name, Root: i.root.ID(), Storage: bytes})
}

// UnmarshalJSON restores memory instance from json representation
func (i *Instance) UnmarshalJSON(b []byte) error {
	var j instanceJSON
	err := json.Unmarshal(b, &j)

	if err != nil {
		return errors.Wrap(err, "failed to unmarshal memory")
	}

	i.name = j.Name

	s, err := i.unmarshalStorageJSONFunc(j.Storage)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal memory storage")
	}

	i.storage = s

	u, err := i.storage.LoadUnit(j.Root)
	if err != nil {
		return errors.Wrap(err, "failed to load root unit from storage")
	}

	i.root = u
	return nil
}
