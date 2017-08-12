package operations

import (
	"fmt"
	"chip8/system"
)

type BinaryCodedDecimalParser struct {}
func(p BinaryCodedDecimalParser) Matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0xF && byte(opcode) == 0x33
}

func(p BinaryCodedDecimalParser) CreateOp(opcode system.OpCode) Operation {
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

func(o BinaryCodedDecimalOp) Execute(machine *system.VirtualMachine) {
	val := machine.Registers[o.register]

	machine.Memory[machine.IndexRegister] = val / 100
	machine.Memory[machine.IndexRegister + 1] = (val / 10) % 10
	machine.Memory[machine.IndexRegister + 2] = val % 10
}
