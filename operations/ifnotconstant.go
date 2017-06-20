package operations

import (
	"fmt"
	"chip8/system"
)

type IfNotConstantParser struct {}
func(incp IfNotConstantParser) Matches(opcode OpCode) bool {
	// TODO:  There is probably a more efficient way to do this
	return opcode.String()[0] == '4'
}

func(incp IfNotConstantParser) CreateOp(opcode OpCode) Operation {
	return IfNotConstantOp{
		register: int8((int16(opcode) & 0x0F00) >> 8),
		value: int8(opcode & 0x00FF),
	}
}

type IfNotConstantOp struct {
	register int8
	value int8
}
func(o IfNotConstantOp) String() string {
	return fmt.Sprintf("If V%X != %X", o.register, o.value)
}

func(o IfNotConstantOp) Execute(machine *system.VirtualMachine) {
	// TODO:  Get this going
}
