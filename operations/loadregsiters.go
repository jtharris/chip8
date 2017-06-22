package operations

import (
	"fmt"
	"chip8/system"
)

type LoadRegistersParser struct {}
func(p LoadRegistersParser) Matches(opcode OpCode) bool {
	return opcode.String()[0] == 'f' && byte(opcode) == 0x65
}

func(p LoadRegistersParser) CreateOp(opcode OpCode) Operation {
	return LoadRegistersOp{
		topRegister: byte(opcode & 0x0F00 >> 8),
	}
}

type LoadRegistersOp struct {
	topRegister byte
}
func(o LoadRegistersOp) String() string {
	return fmt.Sprintf("loadRegisters(V%X, &I)", o.topRegister)
}

func(o LoadRegistersOp) Execute(machine *system.VirtualMachine) {
	// TODO:  Get this going
}
