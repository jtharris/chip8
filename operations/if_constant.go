package operations

import (
	"fmt"
	"chip8/system"
)

type IfConstantParser struct {}
func(p IfConstantParser) Matches(opcode OpCode) bool {
	return opcode.String()[0] == '3'
}

func(p IfConstantParser) CreateOp(opcode OpCode) Operation {
	return IfConstantOp{
		register: uint8(opcode & 0x0F00 >> 8),
		value: uint8(opcode & 0x00FF),
	}
}

type IfConstantOp struct {
	register uint8
	value uint8
}
func(o IfConstantOp) String() string {
	return fmt.Sprintf("If V%X == %X", o.register, o.value)
}

func(o IfConstantOp) Execute(machine *system.VirtualMachine) {
	if (machine.Registers[o.register] == o.value) {
		machine.ProgramCounter++
	}
}
