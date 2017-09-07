package operations

import (
	"fmt"
	"github.com/jtharris/chip8/system"
)

// Parser for DelayTimerOp
type delayTimerParser struct{}

func (p delayTimerParser) matches(opcode system.OpCode) bool {
	return opcode>>12 == 0xF && byte(opcode) == 0x15
}

func (p delayTimerParser) createOp(opcode system.OpCode) Operation {
	return DelayTimerOp{
		register: byte(opcode & 0x0F00 >> 8),
	}
}

// DelayTimerOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#Fx15
type DelayTimerOp struct {
	register byte
}

// String returns a text representation of this operation
func (o DelayTimerOp) String() string {
	return fmt.Sprintf("DT = V%X", o.register)
}

// Execute this operation on the given virtual machine
func (o DelayTimerOp) Execute(vm *system.VirtualMachine) {
	vm.DelayTimer = vm.Registers[o.register]
}
