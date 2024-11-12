package trial_new

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// / GENERIC TYPE
func InputUser[T any](param T) T {
	return param
}

func TestGeneric(t *testing.T) {
	x := InputUser[string]("hahaha")
	assert.Equal(t, x, "hahaha")
}

// / MULTIPLE GENERIC
func Pair[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1, param2)
}

func TestMultipleParam(t *testing.T) {
	Pair("hahah", 32)
	Pair[int, string](32, "hahah")
}

// / COMPARABLE TYPE, hal ini yg membuat tipe generic bisa di bandingkan
func EqualKah[T comparable](value1 T, value2 T) bool {
	if value1 != value2 {
		return false
	}
	return true
}

func TestComparable(t *testing.T) {
	assert.Equal(t, true, EqualKah[int](110, 110))
	assert.Equal(t, true, EqualKah[string]("haha", "haha"))
}
