package v2

import "fmt"

// Key is an interface that defines a method to get a resource key.
type Key interface {
	getResourceKey() string
}

type NameKey struct {
	Name string
}

type IDKey struct {
	ID int
}

// getResourceKey returns the resource key for NameKey.
// It returns the name of the resource.
// getResourceKey returns the resource key for IDKey.
func (k NameKey) getResourceKey() string {
	return k.Name
}

// getResourceKey returns the resource key for IDKey.
// This is a string representation of the ID.
// It is used for generating the URL for the resource.
func (k IDKey) getResourceKey() string {
	return fmt.Sprintf("%d", k.ID)
}
