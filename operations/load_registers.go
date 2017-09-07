package operations

import (
	"github.com/jtharris/chip8/system"
	"fmt"
)

// Parser for LoadRegistersOp
type loadRegistersParser struct{}

func (p loadRegistersParser) matches(opcode system.OpCode) bool {
	return opcode>>12 == 0xF && byte(opcode) == 0x65
}

func (p loadRegistersParser) createOp(opcode system.OpCode) Operation {
	return LoadRegistersOp{
		topRegister: byte(opcode & 0x0F00 >> 8),
	}
}

// LoadRegistersOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#Fx65
type LoadRegistersOp struct {
	topRegister byte
}

// String returns a text representation of this operation
func (o LoadRegistersOp) String() string {
	return fmt.Sprintf("load_registers(V%X, &I)", o.topRegister)
}

// Execute this operation on the given virtual machine
func (o LoadRegistersOp) Execute(vm *system.VirtualMachine) {
	for i := byte(0); i <= o.topRegister; i++ {
		vm.Registers[i] = vm.Memory[vm.IndexRegister+uint16(i)]
	}
}
