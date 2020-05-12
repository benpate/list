package listutil

import "testing"

func TestFirst(t *testing.T) {

	assert.Equal(t, "Hello", First("Hello There", " "))

	assert.Equal(t, "Hello", First("Hello ", " "))

	assert.Equal(t, "Hello", First("Hello", " "))
}
