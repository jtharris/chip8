package operations

import (
	"chip8/system"
	"fmt"
)

// Parser for AddRegisterOp
type addRegisterParser struct{}

func (p addRegisterParser) matches(opcode system.OpCode) bool {
	return opcode>>12 == 0x8 && opcode&0x000F == 0x4
}

func (p addRegisterParser) createOp(opcode system.OpCode) Operation {
	return AddRegisterOp{
		register1: byte(opcode & 0x0F00 >> 8),
		register2: byte(opcode & 0x00F0 >> 4),
	}
}

// AddRegisterOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#8xy4
type AddRegisterOp struct {
	register1 byte
	register2 byte
}

// String returns a text representation of this operation
func (o AddRegisterOp) String() string {
	return fmt.Sprintf("V%X += V%X", o.register1, o.register2)
}

// Execute this operation on the given virtual machine
func (o AddRegisterOp) Execute(vm *system.VirtualMachine) {
	// Set the overflow register first
	vm.Registers[0xF] = byte((uint16(vm.Registers[o.register1]) + uint16(vm.Registers[o.register2])) >> 8)

	// Then add the value
	vm.Registers[o.register1] += vm.Registers[o.register2]
}
