package operations

import (
	"fmt"
	"chip8/system"
)

type AddConstantParser struct {}
func(acp AddConstantParser) Matches(opcode OpCode) bool {
	// TODO:  There is probably a more efficient way to do this
	return opcode.String()[0] == '7'
}

func(acp AddConstantParser) CreateOp(opcode OpCode) Operation {
	return AddConstantOp{
		register: byte(opcode & 0x0F00 >> 8),
		value: byte(opcode & 0x00FF),
	}
}

type AddConstantOp struct {
	register byte
	value byte
}
func(o AddConstantOp) String() string {
	return fmt.Sprintf("Adding V%X += %X", o.register, o.value)
}

func(o AddConstantOp) Execute(machine *system.VirtualMachine) {
	// TODO:  Get this going
}
