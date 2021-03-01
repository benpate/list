package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHead(t *testing.T) {

	assert.Equal(t, "Hello", Head("Hello There", " "))

	assert.Equal(t, "", Head(" Hello", " "))

	assert.Equal(t, "Hello", Head("Hello ", " "))

	assert.Equal(t, "Hello", Head("Hello", " "))

	assert.Equal(t, "This is a very long string but still not a list.", Head("This is a very long string but still not a list.", ","))

	assert.Equal(t, "One,Two,Three", Head("One,Two,Three", " "))
}

func TestTail(t *testing.T) {

	assert.Equal(t, "There", Tail("Hello There", " "))

	assert.Equal(t, "Bananas,Pears", Tail("Apples,Bananas,Pears", ","))

	assert.Equal(t, "", Tail("One,Two,Three", " "))

	assert.Equal(t, "Hello", Tail(" Hello", " "))

	assert.Equal(t, "", Tail("Hello ", " "))

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

func TestRemoveLast(t *testing.T) {

	assert.Equal(t, "Hello There", RemoveLast("Hello There Dude", " "))

	assert.Equal(t, "", RemoveLast(" Hello", " "))

	assert.Equal(t, "Hello", RemoveLast("Hello ", " "))

	assert.Equal(t, "", RemoveLast("Hello", " "))

	assert.Equal(t, "", RemoveLast("This is a very long string but still not a list.", ","))

	assert.Equal(t, "", RemoveLast("One,Two,Three", " "))

	assert.Equal(t, "Hello", RemoveLast("Hello There", " "))

	assert.Equal(t, "Apples,Bananas", RemoveLast("Apples,Bananas,Pears", ","))

	assert.Equal(t, "", RemoveLast("One,Two,Three", " "))

	assert.Equal(t, "", RemoveLast(" Hello", " "))

	assert.Equal(t, "Hello", RemoveLast("Hello ", " "))

}

func TestSplit(t *testing.T) {

	{
		a, b := Split("one,two,three", ",")
		assert.Equal(t, "one", a)
		assert.Equal(t, "two,three", b)
	}

	{
		a, b := Split("one,two,three", "!")
		assert.Equal(t, "one,two,three", a)
		assert.Equal(t, "", b)
	}
}

func TestSplitTail(t *testing.T) {

	{
		a, b := SplitTail("one,two,three", ",")
		assert.Equal(t, "one,two", a)
		assert.Equal(t, "three", b)
	}

	{
		a, b := SplitTail("one,two,three", "!")
		assert.Equal(t, "one,two,three", a)
		assert.Equal(t, "", b)
	}
}

func TestAt(t *testing.T) {

	{
		assert.Equal(t, "one", At("one,two,three", ",", 0))
		assert.Equal(t, "two", At("one,two,three", ",", 1))
		assert.Equal(t, "three", At("one,two,three", ",", 2))
		assert.Equal(t, "", At("one,two,three", ",", 3))
	}
}
