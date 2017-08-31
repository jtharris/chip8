package operations

import (
	"chip8/system"
	"fmt"
)

// Parser for ReverseSubtractRegisterOp
type reverseSubtractRegisterParser struct{}

func (p reverseSubtractRegisterParser) matches(opcode system.OpCode) bool {
	return opcode>>12 == 0x8 && opcode&0x000F == 0x7
}

func (p reverseSubtractRegisterParser) createOp(opcode system.OpCode) Operation {
	return ReverseSubtractRegisterOp{
		register1: uint8(opcode & 0x0F00 >> 8),
		register2: uint8(opcode & 0x00F0 >> 4),
	}
}

// ReverseSubtractRegisterOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#8xy7
type ReverseSubtractRegisterOp struct {
	register1 uint8
	register2 uint8
}

// String returns a text representation of this operation
func (o ReverseSubtractRegisterOp) String() string {
	return fmt.Sprintf("V%X = V%X - V%X", o.register1, o.register2, o.register1)
}

// Execute this operation on the given virtual machine
func (o ReverseSubtractRegisterOp) Execute(vm *system.VirtualMachine) {
	val1 := vm.Registers[o.register1]
	val2 := vm.Registers[o.register2]

	if val1 > val2 {
		vm.Registers[0xF] = 0x0
	} else {
		vm.Registers[0xF] = 0x1
	}

	vm.Registers[o.register1] = val2 - val1
}
