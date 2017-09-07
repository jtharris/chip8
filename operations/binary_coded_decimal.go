package operations

import (
	"github.com/jtharris/chip8/system"
	"fmt"
)

// Parser for BinaryCodedDecimalOop
type binaryCodedDecimalParser struct{}

func (p binaryCodedDecimalParser) matches(opcode system.OpCode) bool {
	return opcode>>12 == 0xF && byte(opcode) == 0x33
}

func (p binaryCodedDecimalParser) createOp(opcode system.OpCode) Operation {
	return BinaryCodedDecimalOp{
		register: byte(opcode & 0x0F00 >> 8),
	}
}

// BinaryCodedDecimalOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#Fx33
type BinaryCodedDecimalOp struct {
	register byte
}

// String returns a text representation of this operation
func (o BinaryCodedDecimalOp) String() string {
	return fmt.Sprintf("BCD(V%X)", o.register)
}

// Execute this operation on the given virtual machine
func (o BinaryCodedDecimalOp) Execute(vm *system.VirtualMachine) {
	val := vm.Registers[o.register]

	vm.Memory[vm.IndexRegister] = val / 100
	vm.Memory[vm.IndexRegister+1] = (val / 10) % 10
	vm.Memory[vm.IndexRegister+2] = val % 10
}
