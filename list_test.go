package listutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {

	assert.Equal(t, "Hello", First("Hello There", " "))

	assert.Equal(t, "Hello", First("Hello ", " "))

	assert.Equal(t, "Hello", First("Hello", " "))

	assert.Equal(t, "This is a very long string but still not a list.", First("This is a very long string but still not a list.", ","))

	assert.Equal(t, "One,Two,Three", First("One,Two,Three", " "))
}

func TestRest(t *testing.T) {

	assert.Equal(t, "There", Rest("Hello There", " "))

	assert.Equal(t, "Bananas,Pears", Rest("Apples,Bananas,Pears", ","))

	assert.Equal(t, "", Rest("One,Two,Three", " "))
}
