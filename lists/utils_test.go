package lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	list := Create(5)
	assert.Equal(t, list.Len(), 5)
	assert.Equal(t, "1 -> 2 -> 3 -> 4 -> 5", list.String())
}

func TestCreateZero(t *testing.T) {
	list := Create(0)
	assert.Nil(t, list)
}

func TestCreateWithStart(t *testing.T) {
	l1 := CreateWithStart(3, 1)
	l2 := CreateWithStart(3, 4)

	assert.Equal(t, "1 -> 2 -> 3", l1.String())
	assert.Equal(t, "4 -> 5 -> 6", l2.String())
}

func TestCreateWithStartInterval(t *testing.T) {
	l1 := CreateWithStartInterval(3, 5, 2)
	assert.Equal(t, "5 -> 7 -> 9", l1.String())
	l1 = CreateWithStartInterval(5, 23, 5)
	assert.Equal(t, "23 -> 28 -> 33 -> 38 -> 43", l1.String())
}
