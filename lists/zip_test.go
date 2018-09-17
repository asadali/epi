package lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	list := Create(4)
	rev := Reverse(list)
	assert.Equal(t, 4, (*rev).Len())
	assert.Equal(t, "4 -> 3 -> 2 -> 1", (*rev).String())
}

func TestFindMidpoint(t *testing.T) {
	list := Create(4)
	mid := FindMidpoint(list)
	assert.Equal(t, 2, (*mid).Len())
	assert.Equal(t, "3 -> 4", (*mid).String())
}

func TestInterleave(t *testing.T) {
	list1 := Create(2)
	list2 := Create(2)

	Interleave(list1, list2)
	assert.Equal(t, 4, (*list1).Len())
	assert.Equal(t, "1 -> 1 -> 2 -> 2", (*list1).String())

	list1 = Create(3)
	list2 = Create(2)

	Interleave(list1, list2)
	assert.Equal(t, 5, (*list1).Len())
	assert.Equal(t, "1 -> 1 -> 2 -> 2 -> 3", (*list1).String())
}

func TestZip(t *testing.T) {
	list := Create(5)
	Zip(list)

	assert.Equal(t, 5, (*list).Len())
	assert.Equal(t, "1 -> 5 -> 2 -> 4 -> 3", (*list).String())

	list = Create(4)
	Zip(list)

	assert.Equal(t, 4, (*list).Len())
	assert.Equal(t, "1 -> 4 -> 2 -> 3", (*list).String())
}
