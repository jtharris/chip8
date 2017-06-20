package operations

import (
	"chip8/system"
	"encoding/hex"
)

type OpCode uint16
func(o OpCode) String() string {
	bytes := []byte{byte(uint16(o) >> 8), byte(o)}
	return hex.EncodeToString(bytes)
}

type OperationParser interface {
	Matches(opcode OpCode) bool
	CreateOp(opcode OpCode) Operation
}

type Operation interface {
	String() string
	Execute(machine *system.VirtualMachine)
}

func CreateOperation(opCode OpCode) Operation {
	// TODO:  Create a static list of parsers
	parsers := []OperationParser{
		ClearParser{},
		ReturnParser{},
		GotoParser{},
		CallParser{},
		IfConstantParser{},
		IfNotConstantParser{},
	}

	for _, parser := range parsers {
		if parser.Matches(opCode) {
			return parser.CreateOp(opCode)
		}
	}

	return UnknownOp{}
}