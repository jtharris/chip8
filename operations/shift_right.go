package operations

import (
	"fmt"
	"chip8/system"
)

// Parser for ShiftRightOp
type shiftRightParser struct {}
func(p shiftRightParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0x8 && opcode & 0x000F == 0x6
}

func(p shiftRightParser) createOp(opcode system.OpCode) Operation {
	return ShiftRightOp{
		register: byte(opcode & 0x0F00 >> 8),
	}
}

// ShiftRightOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#8xy6
type ShiftRightOp struct {
	register byte
}

// String returns a text representation of this operation
func(o ShiftRightOp) String() string {
	return fmt.Sprintf("V%X >> 1", o.register)
}

// Execute this operation on the given virtual machine
func(o ShiftRightOp) Execute(vm *system.VirtualMachine) {
	val := vm.Registers[o.register]

	vm.Registers[0xF] = val & 0x01
	vm.Registers[o.register] = val >> 1
}
