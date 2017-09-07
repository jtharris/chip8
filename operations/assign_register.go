package operations

import (
	"fmt"
	"github.com/jtharris/chip8/system"
)

// Parser for AssignRegisterOp
type assignRegisterParser struct{}

func (p assignRegisterParser) matches(opcode system.OpCode) bool {
	return opcode>>12 == 0x8 && opcode&0x000F == 0x0
}

func (p assignRegisterParser) createOp(opcode system.OpCode) Operation {
	return AssignRegisterOp{
		register1: byte(opcode & 0x0F00 >> 8),
		register2: byte(opcode & 0x00F0 >> 4),
	}
}

// AssignRegisterOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#8xy0
type AssignRegisterOp struct {
	register1 byte
	register2 byte
}

// String returns a text representation of this operation
func (o AssignRegisterOp) String() string {
	return fmt.Sprintf("V%X = V%X", o.register1, o.register2)
}

// Execute this operation on the given virtual machine
func (o AssignRegisterOp) Execute(vm *system.VirtualMachine) {
	vm.Registers[o.register1] = vm.Registers[o.register2]
}
