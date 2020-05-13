package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {

	assert.Equal(t, "Hello", First("Hello There", " "))

	assert.Equal(t, "", First(" Hello", " "))

	assert.Equal(t, "Hello", First("Hello ", " "))

	assert.Equal(t, "Hello", First("Hello", " "))

	assert.Equal(t, "This is a very long string but still not a list.", First("This is a very long string but still not a list.", ","))

	assert.Equal(t, "One,Two,Three", First("One,Two,Three", " "))
}

func TestRest(t *testing.T) {

	assert.Equal(t, "There", Rest("Hello There", " "))

	assert.Equal(t, "Bananas,Pears", Rest("Apples,Bananas,Pears", ","))

	assert.Equal(t, "", Rest("One,Two,Three", " "))

	assert.Equal(t, "Hello", Rest(" Hello", " "))

	assert.Equal(t, "", Rest("Hello ", " "))

}

func TestLast(t *testing.T) {

	assert.Equal(t, "Dude", Last("Hello There Dude", " "))

	assert.Equal(t, "Hello", Last(" Hello", " "))

	assert.Equal(t, "", Last("Hello ", " "))

	assert.Equal(t, "Hello", Last("Hello", " "))

	assert.Equal(t, "This is a very long string but still not a list.", Last("This is a very long string but still not a list.", ","))

	assert.Equal(t, "One,Two,Three", Last("One,Two,Three", " "))

	assert.Equal(t, "There", Last("Hello There", " "))

	assert.Equal(t, "Pears", Last("Apples,Bananas,Pears", ","))

	assert.Equal(t, "One,Two,Three", Last("One,Two,Three", " "))

	assert.Equal(t, "Hello", Last(" Hello", " "))

	assert.Equal(t, "", Last("Hello ", " "))

}
