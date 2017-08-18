package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateOperationFound(t *testing.T) {
	expected := SoundTimerOp{register: 0xA}

	assert.Equal(t, expected, CreateOperation(0xFA18))
}

func TestCreateOperationFoundInCache(t *testing.T) {
	expected := DelayTimerOp{register: 0xA}

	// Get the cache hot
	CreateOperation(0xFA15)

	assert.Equal(t, expected, CreateOperation(0xFA15))
}

func TestCreateOperationNotFound(t *testing.T) {
	expected := UnknownOp{code: 0x5AB1}

	assert.Equal(t, expected, CreateOperation(0x5AB1))
}
