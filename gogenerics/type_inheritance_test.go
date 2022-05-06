package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](param T) string {
	return param.GetName()
}

// Manager
type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

// Manager End

// VicePresident
type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (m *MyVicePresident) GetName() string {
	return m.Name
}

func (m *MyVicePresident) GetVicePresidentName() string {
	return m.Name
}

// VicePresident End

func TestInheritance(t *testing.T) {
	assert.Equal(t, "Nagachi", GetName[Manager](&MyManager{Name: "Nagachi"}))
	assert.Equal(t, "Rafly", GetName[VicePresident](&MyVicePresident{Name: "Rafly"}))
	assert.Equal(t, "Nagachi", GetName(&MyVicePresident{Name: "Nagachi"}))
}
