package trial_new

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/// Type Set,kumpullan tipe data yg memenuhi kriteria atau batasan tettentu (pembatasan tipe generic)
//hanya bisa digunakan sebagai parameter

// CONTOH :
type Number interface {
	~int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func Min[T Number](first T, sec T) T {
	if first < sec {
		return first
	}
	return sec
}

func TestTypeSet(t *testing.T) {
	assert.Equal(t, 3, Min[int](10, 3))

	//TYPE INFERENCE TANPA PENYEBUTAN [] OTOMATIS DETECT
	assert.Equal(t, 3, Min(10, 3))
	assert.Equal(t, 0.4, Min(0.4, 1))
	assert.Equal(t, 7.1, Min(10.0, 7.1))
}

/// type approximation, memungkinakan dalam tipe sets variabel yg puya tipe data yg sama di dalam type sets
// bisa mengunakan metod yg sama dengan simbol (~)

type Age int

func TestApproximation(t *testing.T) {
	assert.Equal(t, Age(10), Min(Age(121), Age(10)))
}
