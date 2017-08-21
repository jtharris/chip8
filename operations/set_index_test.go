package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"chip8/system"
)

func TestSetIndexParser_Matches(t *testing.T) {
	parser := setIndexParser{}

	assert.True(t, parser.matches(0xA8C1))
}

func TestSetIndexParser_DoesNotMatch(t *testing.T) {
	parser := setIndexParser{}

	assert.False(t, parser.matches(0x98C1))
}

func TestSetIndexParser_CreateOp(t *testing.T) {
	parser := setIndexParser{}
	expected := SetIndexOp{value: 0x08C1}

	assert.Equal(t, expected, parser.createOp(0xA8C1))
}

func TestSetIndexOp_String(t *testing.T) {
	op := SetIndexOp{value: 0x090D}

	assert.Equal(t, "I = 90D", op.String())
}

func TestSetIndexOp_Execute(t *testing.T) {
	vm := system.VirtualMachine{}
	op := SetIndexOp{value: 0x090D}

	op.Execute(&vm)

	assert.Equal(t, uint16(0x090D), vm.IndexRegister)
}
