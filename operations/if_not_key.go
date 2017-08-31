package operations

import (
	"fmt"
	"chip8/system"
)

// Parser for IfNotKeyOp
type ifNotKeyParser struct {}
func(p ifNotKeyParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0xE && byte(opcode) == 0xA1
}

func(p ifNotKeyParser) createOp(opcode system.OpCode) Operation {
	return IfNotKeyOp{
		register: byte(opcode & 0x0F00 >> 8),
	}
}

// IfNotKeyOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#ExA1
type IfNotKeyOp struct {
	register byte
}

// String returns a text representation of this operation
func(o IfNotKeyOp) String() string {
	return fmt.Sprintf("If key != V%X", o.register)
}

// Execute this operation on the given virtual machine
func(o IfNotKeyOp) Execute(vm *system.VirtualMachine) {
	key := vm.Registers[o.register]
	if !vm.Keyboard[key] {
		vm.IncrementPC()
	}

	// Clear the keyboard now that the key has been registered.
	// This is needed for inputs that don't have key up events like termbox
	vm.Keyboard[key] = false
}
