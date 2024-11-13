package trial_new

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Max[T interface {
	int | int64 | float32 | float64
}](v1 T, v2 T) T {
	if v1 > v2 {
		return v1
	}
	return v2
}

func TestMax(t *testing.T) {
	assert.Equal(t, 21.9, Max(21.9, 10))
	assert.Equal(t, 321, Max(31, 321))
	assert.Equal(t, 43.21, Max(10.9, 43.21))
}
