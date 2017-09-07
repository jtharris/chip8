package operations

import (
	"fmt"
	"github.com/jtharris/chip8/system"
)

// Parser for GotoOp
type gotoParser struct{}

func (p gotoParser) matches(opcode system.OpCode) bool {
	return opcode>>12 == 0x1
}

func (p gotoParser) createOp(opcode system.OpCode) Operation {
	return GotoOp{
		address: uint16(opcode) & 0x0FFF,
	}
}

// GotoOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#1nnn
type GotoOp struct {
	address uint16
}

// String returns a text representation of this operation
func (o GotoOp) String() string {
	return fmt.Sprintf("Goto: %X", o.address)
}

// Execute this operation on the given virtual machine
func (o GotoOp) Execute(vm *system.VirtualMachine) {
	// The main loop will increment this back to the address after this has executed
	vm.ProgramCounter = o.address - 2
}
