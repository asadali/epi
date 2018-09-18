package lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLen(t *testing.T) {
	list := Create(10)
	assert.Equal(t, 10, list.Len())
}

func TestString(t *testing.T) {
	n := Node{Data: 1}
	assert.Equal(t, n.String(), "1")
}

func TestStringLong(t *testing.T) {
	n1 := Node{Data: 1}
	n2 := Node{Data: 2, Next: &n1}
	assert.Equal(t, n2.String(), "2 -> 1")
}
