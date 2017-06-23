package operations

import (
	"fmt"
	"chip8/system"
)

type IfNotConstantParser struct {}
func(p IfNotConstantParser) Matches(opcode OpCode) bool {
	// TODO:  There is probably a more efficient way to do this
	return opcode.String()[0] == '4'
}

func(p IfNotConstantParser) CreateOp(opcode OpCode) Operation {
	return IfNotConstantOp{
		register: uint8(opcode & 0x0F00 >> 8),
		value: uint8(opcode & 0x00FF),
	}
}

type IfNotConstantOp struct {
	register uint8
	value uint8
}
func(o IfNotConstantOp) String() string {
	return fmt.Sprintf("If V%X != %X", o.register, o.value)
}

func(o IfNotConstantOp) Execute(machine *system.VirtualMachine) {
	// TODO:  Get this going
}
