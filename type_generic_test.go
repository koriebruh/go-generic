package trial_new

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestGenericType(t *testing.T) {
	insideBag := Bag[string]{"ruler", "eraser", "book", "phone"}
	PrintBag(insideBag)

	bagNumber := Bag[int]{1, 2, 34, 5, 6, 6732, 21}
	PrintBag(bagNumber)
}
