package operations

import (
	"fmt"
	"chip8/system"
)

type SetToDelayParser struct {}
func(p SetToDelayParser) Matches(opcode OpCode) bool {
	return opcode.String()[0] == 'f' && byte(opcode) == 0x07
}

func(p SetToDelayParser) CreateOp(opcode OpCode) Operation {
	return SetToDelayOp{
		register: byte(opcode & 0x0F00 >> 8),
	}
}

type SetToDelayOp struct {
	register byte
}
func(o SetToDelayOp) String() string {
	return fmt.Sprintf("V%X = delay_timer", o.register)
}

func(o SetToDelayOp) Execute(machine *system.VirtualMachine) {
	machine.Registers[o.register] = machine.DelayTimer
}
