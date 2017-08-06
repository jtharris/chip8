package operations

import (
	"chip8/system"
	"fmt"
)


type DrawParser struct {}
func(p DrawParser) Matches(opcode system.OpCode) bool {
	return opcode.String()[0] == 'd'
}

func(p DrawParser) CreateOp(opcode system.OpCode) Operation {
	return DrawOp{
		xRegister: byte(opcode & 0x0F00 >> 8),
		yRegister: byte(opcode & 0x00F0 >> 4),
		height: byte(opcode & 0x000F),
	}
}

type DrawOp struct {
	xRegister byte
	yRegister byte
	height byte
}

func(o DrawOp) String() string {
	return fmt.Sprintf("Draw Screen (V%X, V%X) Height: %X", o.xRegister, o.yRegister, o.height)
}

func(o DrawOp) Execute(machine *system.VirtualMachine) {
	machine.Registers[0xF] = 0	// start with this as the default position
	xPos := machine.Registers[o.xRegister]
	yPos := machine.Registers[o.yRegister]

	for row := byte(0); row < o.height; row++ {
		y := yPos + row

		// Edge case where sprites can be drawn off the screen?
		if int(y) >= len(machine.Pixels) {
			return
		}

		sprite := uint64(machine.Memory[machine.IndexRegister + uint16(row)])
		offset := 56 - int(xPos)

		if offset > 0 {
			sprite = sprite << uint(offset)
		} else {
			sprite = sprite >> uint(-offset)
		}

		// If any 'on' pixels are going to be flipped, then set
		// VF to 1 per the spec
		if sprite & machine.Pixels[y] > 0 {
			machine.Registers[0xF] = 1
		}

		machine.Pixels[y] ^= sprite
	}
}
