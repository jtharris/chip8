package operations

import (
	"github.com/jtharris/chip8/system"
	"fmt"
)

// Parser for SetIndexOp
type setIndexParser struct{}

func (p setIndexParser) matches(opcode system.OpCode) bool {
	return opcode>>12 == 0xA
}

func (p setIndexParser) createOp(opcode system.OpCode) Operation {
	return SetIndexOp{
		value: uint16(opcode) & 0x0FFF,
	}
}

// SetIndexOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#Annn
type SetIndexOp struct {
	value uint16
}

// String returns a text representation of this operation
func (o SetIndexOp) String() string {
	return fmt.Sprintf("I = %X", o.value)
}

// Execute this operation on the given virtual machine
func (o SetIndexOp) Execute(vm *system.VirtualMachine) {
	vm.IndexRegister = o.value
}
