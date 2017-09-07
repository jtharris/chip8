package operations

import (
	"github.com/jtharris/chip8/system"
	"fmt"
	"math/rand"
)

// Parser for RandomOp
type randomParser struct{}

func (p randomParser) matches(opcode system.OpCode) bool {
	return opcode>>12 == 0xC
}

func (p randomParser) createOp(opcode system.OpCode) Operation {
	return RandomOp{
		register: byte(opcode & 0x0F00 >> 8),
		value:    byte(opcode),
	}
}

// RandomOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#Cxkk
type RandomOp struct {
	register byte
	value    byte
}

// String returns a text representation of this operation
func (o RandomOp) String() string {
	return fmt.Sprintf("V%X = rand(255) & %X", o.register, o.value)
}

// Execute this operation on the given virtual machine
func (o RandomOp) Execute(vm *system.VirtualMachine) {
	vm.Registers[o.register] = byte(rand.Int()) & o.value
}
