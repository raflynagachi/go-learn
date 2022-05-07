package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindMin[T interface{ int | float64 }](value1, value2 T) T {
	if value1 <= value2 {
		return value1
	} else {
		return value2
	}
}

func GetFirst[T []E, E any](slice T) E {
	first := slice[0]
	return first
}

func TestInlineConstraint(t *testing.T) {
	assert.Equal(t, 4, FindMin(4, 7))
	assert.Equal(t, 4.0, FindMin(4.0, 15.00))

	names := []string{
		"Rafly", "Rigan", "Nagachi",
	}
	assert.Equal(t, "Rafly", GetFirst[[]string, string](names))
}
