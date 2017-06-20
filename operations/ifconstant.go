package operations

import (
	"fmt"
	"chip8/system"
)

type IfConstantParser struct {}
func(icp IfConstantParser) Matches(opcode OpCode) bool {
	// TODO:  There is probably a more efficient way to do this
	return opcode.String()[0] == '3'
}

func(icp IfConstantParser) CreateOp(opcode OpCode) Operation {
	return IfConstantOp{
		register: int8((int16(opcode) & 0x0F00) >> 8),
		value: int8(opcode & 0x00FF),
	}
}

type IfConstantOp struct {
	register int8
	value int8
}
func(o IfConstantOp) String() string {
	return fmt.Sprintf("If V%X == %X", o.register, o.value)
}

func(o IfConstantOp) Execute(machine *system.VirtualMachine) {
	// TODO:  Get this going
}
