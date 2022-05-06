package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

type Number interface {
	// type approximation: all aliases of int
	// are compatible
	~int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func Min[T Number](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestTypeSet(t *testing.T) {
	assert.Equal(t, 100, Min[int](100.0, 300))
	assert.Equal(t, 4.0, Min[float64](32, 4))

	//Age is alias of int
	assert.Equal(t, 4.0, Min[Age](32, 4))
}
