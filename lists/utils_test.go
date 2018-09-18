package lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	list := Create(5)
	assert.Equal(t, list.Len(), 5)
	assert.Equal(t, list.String(), "1 -> 2 -> 3 -> 4 -> 5")
}

func TestCreateZero(t *testing.T) {
	list := Create(0)
	assert.Nil(t, list)
}
