package operations

import (
	"fmt"
	"chip8/system"
)

type BinaryCodedDecimalParser struct {}
func(p BinaryCodedDecimalParser) Matches(opcode OpCode) bool {
	return opcode.String()[0] == 'f' && byte(opcode) == 0x33
}

func(p BinaryCodedDecimalParser) CreateOp(opcode OpCode) Operation {
	return BinaryCodedDecimalOp{
		register: byte(opcode & 0x0F00 >> 8),
	}
}

type BinaryCodedDecimalOp struct {
	register byte
}
func(o BinaryCodedDecimalOp) String() string {
	return fmt.Sprintf("Set BCD(V%X)", o.register)
}

func(o BinaryCodedDecimalOp) Execute(machine *system.VirtualMachine) {
	// TODO:  Get this going
}
