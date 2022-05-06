package gogenerics

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[V any](bag Bag[V]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestGenericType(t *testing.T) {
	t.Run("BagString", func(t *testing.T) {
		names := Bag[string]{"Rafly", "Rigan", "Nagachi"}
		PrintBag(names)
	})
	t.Run("BagInt", func(t *testing.T) {
		numbers := Bag[int]{1, 2, 3, 4}
		PrintBag(numbers)
	})
}
