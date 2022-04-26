package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := InitializeService(true)
	assert.NotNil(t, err)
	assert.Nil(t, simpleService)
}

func TestSimpleServiceSuccess(t *testing.T) {
	simplService, err := InitializeService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simplService)
}
