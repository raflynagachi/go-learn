package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[E]) ChangeFirst(first E) E {
	d.First = first
	return d.First
}

func TestGenericStruct(t *testing.T) {
	data := Data[string]{
		First:  "Rafly",
		Second: "Nagachi",
	}
	data2 := Data[int]{
		First:  21,
		Second: 12,
	}
	assert.Equal(t, "Nagachi", data.Second)
	assert.Equal(t, 21, data2.First)

	assert.Equal(t, "Rigan", data.ChangeFirst("Rigan"))
	assert.Equal(t, "Dragon", data.ChangeFirst("Dragon"))
}
