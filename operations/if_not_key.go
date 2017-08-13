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
		register: byte(opcode & 0x0F00 >> 8),
	}
}

type IfNotKeyOp struct {
	register byte
}
func(o IfNotKeyOp) String() string {
	return fmt.Sprintf("If key != V%X", o.register)
}

func(o IfNotKeyOp) Execute(machine *system.VirtualMachine) {
	key := machine.Registers[o.register]
	if !machine.Keyboard[key] {
		machine.IncrementPC()
	}

	// Clear the keyboard now that the key has been registered.
	// This is needed for inputs that don't have key up events like termbox
	machine.Keyboard[key] = false
}
