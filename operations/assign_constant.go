package operations

import (
	"fmt"
	"chip8/system"
)

type AssignConstantParser struct {}
func(p AssignConstantParser) Matches(opcode system.OpCode) bool {
	return opcode.String()[0] == '6'
}

func(p AssignConstantParser) CreateOp(opcode system.OpCode) Operation {
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
	return fmt.Sprintf("V%X = %X", o.register, o.value)
}

func(o AssignConstantOp) Execute(machine *system.VirtualMachine) {
	machine.Registers[o.register] = o.value
}
