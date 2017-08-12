package operations

import (
	"fmt"
	"chip8/system"
)

type IfNotKeyParser struct {}
func(p IfNotKeyParser) Matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0xE && byte(opcode) == 0xA1
}

func(p IfNotKeyParser) CreateOp(opcode system.OpCode) Operation {
	return IfNotKeyOp{
		register: uint8(opcode & 0x0F00 >> 8),
	}
}

type IfNotKeyOp struct {
	register uint8
}
func(o IfNotKeyOp) String() string {
	return fmt.Sprintf("If key != V%X", o.register)
}

func(o IfNotKeyOp) Execute(machine *system.VirtualMachine) {
	if !machine.Keyboard[machine.Registers[o.register]] {
		// TODO:  Move this
		machine.ProgramCounter += 2
	}
}
