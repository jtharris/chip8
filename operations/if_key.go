package operations

import (
	"github.com/jtharris/chip8/system"
	"fmt"
)

// Parser for IfKeyOp
type ifKeyParser struct{}

func (p ifKeyParser) matches(opcode system.OpCode) bool {
	return opcode>>12 == 0xE && byte(opcode) == 0x9E
}

func (p ifKeyParser) createOp(opcode system.OpCode) Operation {
	return IfKeyOp{
		register: byte(opcode & 0x0F00 >> 8),
	}
}

// IfKeyOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#Ex9E
type IfKeyOp struct {
	register byte
}

// String returns a text representation of this operation
func (o IfKeyOp) String() string {
	return fmt.Sprintf("If key == V%X", o.register)
}

// Execute this operation on the given virtual machine
func (o IfKeyOp) Execute(vm *system.VirtualMachine) {
	key := vm.Registers[o.register]
	if vm.Keyboard[key] {
		vm.IncrementPC()
	}

	// Clear the keyboard now that the key has been registered.
	// This is needed for inputs that don't have key up events like termbox
	vm.Keyboard[key] = false
}
