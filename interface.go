package memories

import (
	"github.com/anothermemory/memory"
)

// Interface represents some collection of memory instances
type Interface interface {
	// Add adds new memory to the collection with the given name
	Add(name string, memory memory.Interface) error

	// Remove removes memory with given name from the collection
	Remove(name string) error

	// Get returns memory with given name from the collection
	Get(name string) (memory.Interface, error)

	// GetAll returns map for memories and their names from the collection
	GetAll() (map[string]memory.Interface, error)

	// RemoveAll removes all memories from the collection
	RemoveAll() error
}
