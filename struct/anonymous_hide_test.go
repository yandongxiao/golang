// What are the rules when there are two fields with the same name
// An outer name hides an inner name.
// This provides a way to override(重载) a field or method.
//
// NOTE : If the same name appears twice at the same level, it is an error
// if the name is used by the program. If it’s not used, it doesn’t matter.
// There are no rules to resolve the ambiguity; it must be fixed.
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Animal struct {
	Name string
}

type Cat struct {
	Name string
	Animal
}

func TestHide(t *testing.T) {
	cat := Cat{
		Name:   "Cat",
		Animal: Animal{Name: "Animal"},
	}
	assert.True(t, cat.Name == "Cat")
}
