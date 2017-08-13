package operations

import (
	"fmt"
	"chip8/system"
)

type IfKeyParser struct {}
func(p IfKeyParser) Matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0xE && byte(opcode) == 0x9E
}

func(p IfKeyParser) CreateOp(opcode system.OpCode) Operation {
	return IfKeyOp{
		register: uint8(opcode & 0x0F00 >> 8),
	}
}

type IfKeyOp struct {
	register uint8
}
func(o IfKeyOp) String() string {
	return fmt.Sprintf("If key == V%X", o.register)
}

func(o IfKeyOp) Execute(machine *system.VirtualMachine) {
	if machine.Keyboard[machine.Registers[o.register]] {
		// TODO:  Move this
		machine.ProgramCounter += 2
	}

	// Clear the keyboard now that the key has been registered.
	// This is needed for inputs that don't have key up events like termbox
	machine.Keyboard[machine.Registers[o.register]] = false
}
