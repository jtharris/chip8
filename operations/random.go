package operations

import (
	"fmt"
	"chip8/system"
)

type RandomParser struct {}
func(p RandomParser) Matches(opcode OpCode) bool {
	return opcode.String()[0] == 'c'
}

func(p RandomParser) CreateOp(opcode OpCode) Operation {
	return RandomOp{
		register: byte(opcode & 0x0F00 >> 8),
		value: byte(opcode),
	}
}

type RandomOp struct {
	register byte
	value byte
}
func(o RandomOp) String() string {
	return fmt.Sprintf("V%X = rand(255) & %X", o.register, o.value)
}

func(o RandomOp) Execute(machine *system.VirtualMachine) {
	// TODO:  Get this going
}
