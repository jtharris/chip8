package operations

import (
	"chip8/system"
	"fmt"
)

// Parser for SetToDelayOp
type setToDelayParser struct{}

func (p setToDelayParser) matches(opcode system.OpCode) bool {
	return opcode>>12 == 0xF && byte(opcode) == 0x07
}

func (p setToDelayParser) createOp(opcode system.OpCode) Operation {
	return SetToDelayOp{
		register: byte(opcode & 0x0F00 >> 8),
	}
}

// SetToDelayOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#Fx07
type SetToDelayOp struct {
	register byte
}

// String returns a text representation of this operation
func (o SetToDelayOp) String() string {
	return fmt.Sprintf("V%X = DT", o.register)
}

// Execute this operation on the given virtual machine
func (o SetToDelayOp) Execute(vm *system.VirtualMachine) {
	vm.Registers[o.register] = vm.DelayTimer
}
