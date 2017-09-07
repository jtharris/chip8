package operations

import (
	"github.com/jtharris/chip8/system"
	"fmt"
)

// Parser for ShiftLeftOp
type shiftLeftParser struct{}

func (p shiftLeftParser) matches(opcode system.OpCode) bool {
	return opcode>>12 == 0x8 && opcode&0x000F == 0xE
}

func (p shiftLeftParser) createOp(opcode system.OpCode) Operation {
	return ShiftLeftOp{
		register: byte(opcode & 0x0F00 >> 8),
	}
}

// ShiftLeftOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#8xyE
type ShiftLeftOp struct {
	register byte
}

// String returns a text representation of this operation
func (o ShiftLeftOp) String() string {
	return fmt.Sprintf("V%X << 1", o.register)
}

// Execute this operation on the given virtual machine
func (o ShiftLeftOp) Execute(vm *system.VirtualMachine) {
	val := vm.Registers[o.register]

	vm.Registers[0xF] = val >> 7
	vm.Registers[o.register] = val << 1
}
