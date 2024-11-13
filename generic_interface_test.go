package trial_new

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type data[T any] struct {
	value T
}

func (d *data[T]) GetValue() T {
	return d.value
}

func (d *data[T]) SetValue(value T) {
	d.value = value
}

func TestGenericInterface(t *testing.T) {

	myData := data[string]{}
	result := ChangeValue[string](&myData, "wowoww")
	assert.Equal(t, result, myData.GetValue())

	myData.SetValue("asu")
	assert.Equal(t, "asu", myData.GetValue())
}
