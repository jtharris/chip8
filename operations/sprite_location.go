package operations

import (
	"fmt"
	"chip8/system"
)

// Parser for SpriteLocationOp
type spriteLocationParser struct {}
func(p spriteLocationParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0xF && byte(opcode) == 0x29
}

func(p spriteLocationParser) createOp(opcode system.OpCode) Operation {
	return SpriteLocationOp{
		register: byte(opcode & 0x0F00 >> 8),
	}
}

// SpriteLocationOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#Fx29
type SpriteLocationOp struct {
	register byte
}

// String returns a text representation of this operation
func(o SpriteLocationOp) String() string {
	return fmt.Sprintf("I = sprite_address(V%X)", o.register)
}

// Execute this operation on the given virtual machine
func(o SpriteLocationOp) Execute(vm *system.VirtualMachine) {
	// Each character is 5 bytes wide and starts at memory location 0
	vm.IndexRegister = uint16(vm.Registers[o.register]) * 5
}
