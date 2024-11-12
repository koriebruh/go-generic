package trial_new

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Data[T any] struct {
	FirstName T
	LastName  T
}

func TestGenericStruct(t *testing.T) {
	user1 := Data[any]{
		FirstName: "MUHAMMAD",
		LastName:  "JAMALUDIN NUR",
	}

	assert.Equal(t, "MUHAMMAD", user1.FirstName)
	assert.Equal(t, "JAMALUDIN NUR", user1.LastName)
	assert.NotNil(t, user1)
}

/// METHOD, (func yg di miliki struct)

func (d *Data[_]) SayHello(name string) string {
	return "hello " + name
}

func (d *Data[T]) SetFirstName(newFirstName T) T {
	d.FirstName = newFirstName
	return newFirstName
}

func TestGenericMethod(t *testing.T) {
	user1 := Data[any]{
		FirstName: "MUHAMMAD",
		LastName:  "JAMALUDIN NUR",
	}

	assert.Equal(t, "hello eko", user1.SayHello("eko"))
	assert.Equal(t, "KING", user1.SetFirstName("KING"))

}
