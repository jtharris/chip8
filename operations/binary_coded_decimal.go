package operations

import (
	"fmt"
	"chip8/system"
)

type binaryCodedDecimalParser struct {}
func(p binaryCodedDecimalParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0xF && byte(opcode) == 0x33
}

func(p binaryCodedDecimalParser) createOp(opcode system.OpCode) Operation {
	return BinaryCodedDecimalOp{
		register: byte(opcode & 0x0F00 >> 8),
	}
}

type BinaryCodedDecimalOp struct {
	register byte
}
func(o BinaryCodedDecimalOp) String() string {
	return fmt.Sprintf("BCD(V%X)", o.register)
}

func(o BinaryCodedDecimalOp) Execute(vm *system.VirtualMachine) {
	val := vm.Registers[o.register]

	vm.Memory[vm.IndexRegister] = val / 100
	vm.Memory[vm.IndexRegister + 1] = (val / 10) % 10
	vm.Memory[vm.IndexRegister + 2] = val % 10
}
