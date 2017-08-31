package operations

import (
	"fmt"
	"chip8/system"
)

// Parser for DumpRegistersOp
type dumpRegistersParser struct {}
func(p dumpRegistersParser) matches(opcode system.OpCode) bool {
	return opcode >> 12 == 0xF && byte(opcode) == 0x55
}

func(p dumpRegistersParser) createOp(opcode system.OpCode) Operation {
	return DumpRegistersOp{
		topRegister: byte(opcode & 0x0F00 >> 8),
	}
}

// DumpRegistersOp - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#Fx55
type DumpRegistersOp struct {
	topRegister byte
}

// String returns a text representation of this operation
func(o DumpRegistersOp) String() string {
	return fmt.Sprintf("dump_registers(V%X, &I)", o.topRegister)
}

// Execute this operation on the given virtual machine
func(o DumpRegistersOp) Execute(vm *system.VirtualMachine) {
	for i := byte(0); i <= o.topRegister; i++ {
		vm.Memory[vm.IndexRegister + uint16(i)] = vm.Registers[i]
	}
}
