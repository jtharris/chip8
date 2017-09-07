package operations

import (
	"github.com/jtharris/chip8/system"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnknownOp_StringDefault(t *testing.T) {
	op := UnknownOp{}

	assert.Equal(t, "Unknown Operation:  0000", op.String())
}

func TestUnknownOp_String(t *testing.T) {
	op := UnknownOp{code: 0xF8AA}

	assert.Equal(t, "Unknown Operation:  F8AA", op.String())
}

func TestUnknownOp_Execute(t *testing.T) {
	vm := system.VirtualMachine{}
	op := UnknownOp{code: 0x1234}

	assert.PanicsWithValue(t, "Unknown Operation:  1234", func() { op.Execute(&vm) })
}
