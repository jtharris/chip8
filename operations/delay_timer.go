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
		register: byte(opcode & 0x0F00 >> 8),
	}
}

type DelayTimerOp struct {
	register byte
}
func(o DelayTimerOp) String() string {
	return fmt.Sprintf("delay_timer = V%X", o.register)
}

func(o DelayTimerOp) Execute(machine *system.VirtualMachine) {
	machine.DelayTimer = machine.Registers[o.register]
}
