package lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	l1 := CreateWithStart(3, 1)
	l2 := CreateWithStart(3, 4)

	res := MergeSorted(l1, l2)
	assert.Equal(t, "1 -> 2 -> 3 -> 4 -> 5 -> 6", res.String())
}

func TestMergeOverlap(t *testing.T) {
	l1 := CreateWithStartInterval(5, 23, 2)
	l2 := CreateWithStartInterval(3, 24, 3)

	l1 = MergeSorted(l1, l2)
	assert.Equal(t, 8, l1.Len())
	assert.Equal(t, "23 -> 24 -> 25 -> 27 -> 27 -> 29 -> 30 -> 31", l1.String())

	l2 = CreateWithStartInterval(3, 12, 6)
	l1 = MergeSorted(l1, l2)
	assert.Equal(t, 11, l1.Len())
	assert.Equal(t, "12 -> 18 -> 23 -> 24 -> 24 -> 25 -> 27 -> 27 -> 29 -> 30 -> 31", l1.String())
}
