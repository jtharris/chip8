package operations

import (
	"fmt"
	"chip8/system"
)

type SoundTimerParser struct {}
func(p SoundTimerParser) Matches(opcode OpCode) bool {
	return opcode.String()[0] == 'f' && byte(opcode) == 0x18
}

func(p SoundTimerParser) CreateOp(opcode OpCode) Operation {
	return SoundTimerOp{
		value: byte(opcode & 0x0F00 >> 8),
	}
}

type SoundTimerOp struct {
	value byte
}
func(o SoundTimerOp) String() string {
	return fmt.Sprintf("Set sound timer:  %X", o.value)
}

func(o SoundTimerOp) Execute(machine *system.VirtualMachine) {
	// TODO:  Get this going
}
