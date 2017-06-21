package operations

import (
	"fmt"
	"chip8/system"
)

type AssignRegisterParser struct {}
func(arp AssignRegisterParser) Matches(opcode OpCode) bool {
	// TODO:  There is probably a more efficient way to do this
	opString := opcode.String()
	return opString[0] == '8' && opString[3] == '0'
}

func(arp AssignRegisterParser) CreateOp(opcode OpCode) Operation {
	return AssignRegisterOp{
		register1: byte(opcode & 0x0F00 >> 8),
		register2: byte(opcode & 0x00F0 >> 4),
	}
}

type AssignRegisterOp struct {
	register1 byte
	register2 byte
}
func(o AssignRegisterOp) String() string {
	return fmt.Sprintf("Assign Register V%X = V%X", o.register1, o.register2)
}

func(o AssignRegisterOp) Execute(machine *system.VirtualMachine) {
	// TODO:  Get this going
}
