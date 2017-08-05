package operations

import (
	"fmt"
	"chip8/system"
)

type SpriteLocationParser struct {}
func(p SpriteLocationParser) Matches(opcode system.OpCode) bool {
	return opcode.String()[0] == 'f' && byte(opcode) == 0x29
}

func(p SpriteLocationParser) CreateOp(opcode system.OpCode) Operation {
	return SpriteLocationOp{
		register: byte(opcode & 0x0F00 >> 8),
	}
}

type SpriteLocationOp struct {
	register byte
}
func(o SpriteLocationOp) String() string {
	return fmt.Sprintf("I = sprite_address(V%X)", o.register)
}

func(o SpriteLocationOp) Execute(machine *system.VirtualMachine) {
	// Each character is 5 bytes wide and starts at memory location 0
	machine.IndexRegister = uint16(machine.Registers[o.register]) * 5
}
