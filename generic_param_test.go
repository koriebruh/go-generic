package trial_new

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetFirst[T []E, E any](data T) E {
	first := data[0]
	return first
}

func TestGetFirst(t *testing.T) {
	name := []string{
		"muhammad", "jamaludin", "nur",
	}

	//first := GetFirst[[]string, string](name)
	first := GetFirst(name)
	assert.Equal(t, "muhammad", first)
}
