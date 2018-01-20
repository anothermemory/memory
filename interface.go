package memory

import (
	"encoding/json"

	"github.com/anothermemory/storage"
	"github.com/anothermemory/unit"
)

// Interface represents memory interface
type Interface interface {
	// Storage must be able to marshal it's settings to json
	json.Marshaler

	// Storage must be able to unmarshal it's settings from json
	json.Unmarshaler

	Name() string
	SetName(n string)
	Root() unit.Unit
	SetRoot(u unit.Unit)
	Storage() storage.Interface
}
