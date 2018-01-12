package memory

import (
	"github.com/anothermemory/storage"
	"github.com/anothermemory/unit"
)

type Interface interface {
	Name() string
	SetName(n string)
	Root() unit.Unit
	SetRoot(u unit.Unit)
	Storage() storage.Interface
}
