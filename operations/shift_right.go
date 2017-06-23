package operations

import (
	"fmt"
	"chip8/system"
)

type ShiftRightParser struct {}
func(p ShiftRightParser) Matches(opcode OpCode) bool {
	// TODO:  There is probably a more efficient way to do this
	opString := opcode.String()
	return opString[0] == '8' && opString[3] == '6'
}

func(p ShiftRightParser) CreateOp(opcode OpCode) Operation {
	return ShiftRightOp{
		register: byte(opcode & 0x0F00 >> 8),
	}
}

type ShiftRightOp struct {
	register byte
}
func(o ShiftRightOp) String() string {
	return fmt.Sprintf("V%X >> 1", o.register)
}

func(o ShiftRightOp) Execute(machine *system.VirtualMachine) {
	// TODO:  Get this going
}
