package operations

import (
	"chip8/system"
	"fmt"
)


type drawParser struct {}
func(p drawParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0xD
}

func(p drawParser) createOp(opcode system.OpCode) Operation {
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

func(o DrawOp) Execute(vm *system.VirtualMachine) {
	vm.Registers[0xF] = 0 // start with this as the default position
	xPos := vm.Registers[o.xRegister]
	yPos := vm.Registers[o.yRegister]

	for row := byte(0); row < o.height; row++ {
		y := yPos + row

		// Edge case where sprites can be drawn off the screen?
		if int(y) >= len(vm.Pixels) {
			return
		}

		sprite := uint64(vm.Memory[vm.IndexRegister + uint16(row)])
		offset := 56 - int(xPos)

		if offset > 0 {
			sprite = sprite << uint(offset)
		} else {
			sprite = sprite >> uint(-offset)
		}

		// If any 'on' pixels are going to be flipped, then set
		// VF to 1 per the spec
		if sprite & vm.Pixels[y] > 0 {
			vm.Registers[0xF] = 1
		}

		vm.Pixels[y] ^= sprite
	}
}
