package operations

import (
	"github.com/jtharris/chip8/system"
)

// Parser for ClearOp
type clearParser struct{}

func (p clearParser) matches(opcode system.OpCode) bool {
	return opcode == 0x00E0
}

func (p clearParser) createOp(opcode system.OpCode) Operation {
	return ClearOp{}
}

// ClearOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#00E0
type ClearOp struct{}

// String returns a text representation of this operation
func (o ClearOp) String() string {
	return "Clear Screen"
}

// Execute this operation on the given virtual machine
func (o ClearOp) Execute(vm *system.VirtualMachine) {
	for row := 0; row < len(vm.Pixels); row++ {
		vm.Pixels[row] = 0
	}
}
