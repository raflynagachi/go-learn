package gogenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	var resultString = Length("Nagachi")
	var resultFloat = Length[float64](100)

	assert.Equal(t, "Nagachi", resultString)
	assert.Equal(t, 100.0, resultFloat)
}
