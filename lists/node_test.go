package lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLen(t *testing.T) {
	list := Create(10)
	assert.Equal(t, 10, list.Len())
}
