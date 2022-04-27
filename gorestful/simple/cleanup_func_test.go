package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanupFunc(t *testing.T) {
	file, cleanup, err := InitializeConnection("Nagachi")
	assert.Nil(t, err)
	assert.NotNil(t, file)
	cleanup()
}
