package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

// implement GetterSetter
type MyData[T any] struct{ Value T }

func (m *MyData[T]) GetValue() T {
	return m.Value
}

func (m *MyData[T]) SetValue(value T) {
	m.Value = value
}

func TestGenericInterface(t *testing.T) {
	data := MyData[string]{}
	ChangeValue[string](&data, "Nagachi")
	assert.Equal(t, "Nagachi", data.GetValue())
}
