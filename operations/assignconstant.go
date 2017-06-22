package operations

import (
	"fmt"
	"chip8/system"
)

type AssignConstantParser struct {}
func(p AssignConstantParser) Matches(opcode OpCode) bool {
	// TODO:  There is probably a more efficient way to do this
	return opcode.String()[0] == '6'
}

func(p AssignConstantParser) CreateOp(opcode OpCode) Operation {
	return AssignConstantOp{
		register: byte(opcode & 0x0F00 >> 8),
		value: byte(opcode & 0x00FF),
	}
}

type AssignConstantOp struct {
	register byte
	value byte
}
func(o AssignConstantOp) String() string {
	return fmt.Sprintf("Assign V%X = %X", o.register, o.value)
}

func(o AssignConstantOp) Execute(machine *system.VirtualMachine) {
	// TODO:  Get this going
}
