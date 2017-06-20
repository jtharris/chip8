package operations

import "chip8/system"

type UnknownOp struct{}

func (o UnknownOp) String() string {
	return "Unknown Operation"
}

func (o UnknownOp) Execute(machine *system.VirtualMachine) {
	// Do nothing sentence
}
