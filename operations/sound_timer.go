package operations

import (
	"fmt"
	"github.com/jtharris/chip8/system"
)

// Parser for SoundTimerOp
type soundTimerParser struct{}

func (p soundTimerParser) matches(opcode system.OpCode) bool {
	return opcode>>12 == 0xF && byte(opcode) == 0x18
}

func (p soundTimerParser) createOp(opcode system.OpCode) Operation {
	return SoundTimerOp{
		register: byte(opcode & 0x0F00 >> 8),
	}
}

// SoundTimerOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#Fx18
type SoundTimerOp struct {
	register byte
}

// String returns a text representation of this operation
func (o SoundTimerOp) String() string {
	return fmt.Sprintf("ST = V%X", o.register)
}

// Execute this operation on the given virtual machine
func (o SoundTimerOp) Execute(vm *system.VirtualMachine) {
	vm.SoundTimer = vm.Registers[o.register]
}
