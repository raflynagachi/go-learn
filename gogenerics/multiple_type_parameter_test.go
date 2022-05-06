package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2) (T1, T2) {
	return param1, param2
}

func TestMultipleParameter(t *testing.T) {
	res1, res2 := MultipleParameter[string, int]("Nagachi", 100.0)
	res3, res4 := MultipleParameter(4.0, 100)
	assert.Equal(t, "Nagachi", res1)
	assert.Equal(t, 100, res2)
	assert.Equal(t, 4.0, res3)
	assert.Equal(t, 100, res4)
}

// comparable data type
func isSame[T comparable](num1, num2 T) bool {
	if num1 == num2 {
		return true
	} else {
		return false
	}
}

func TestComparable(t *testing.T) {
	res1 := isSame[float64](4, 2)
	assert.Equal(t, false, res1)

	res2 := isSame[string]("nagachi", "nagachi")
	assert.Equal(t, true, res2)
}
