package operations

import (
	"fmt"
	"chip8/system"
)

type DelayTimerParser struct {}
func(p DelayTimerParser) Matches(opcode OpCode) bool {
	return opcode.String()[0] == 'f' && byte(opcode) == 0x15
}

func(p DelayTimerParser) CreateOp(opcode OpCode) Operation {
	return DelayTimerOp{
		value: byte(opcode & 0x0F00 >> 8),
	}
}

type DelayTimerOp struct {
	value byte
}
func(o DelayTimerOp) String() string {
	return fmt.Sprintf("Set delay timer:  %X", o.value)
}

func(o DelayTimerOp) Execute(machine *system.VirtualMachine) {
	// TODO:  Get this going
}
