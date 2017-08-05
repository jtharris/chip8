package operations

import (
	"chip8/system"
)


type OperationParser interface {
	Matches(opcode system.OpCode) bool
	CreateOp(opcode system.OpCode) Operation
}

type Operation interface {
	String() string
	Execute(machine *system.VirtualMachine)
}

func CreateOperation(opCode system.OpCode) Operation {
	// TODO:  Create a static list of parsers
	parsers := []OperationParser{
		ClearParser{},
		ReturnParser{},
		GotoParser{},
		CallParser{},
		IfConstantParser{},
		IfNotConstantParser{},
		IfRegisterParser{},
		AssignConstantParser{},
		AddConstantParser{},
		AssignRegisterParser{},
		BitwiseOrParser{},
		BitwiseAndParser{},
		BitwiseXorParser{},
		AddRegisterParser{},
		SubtractRegisterParser{},
		ShiftRightParser{},
		ReverseSubtractRegisterParser{},
		SetIndexParser{},
		DrawParser{},
		DelayTimerParser{},
		SetToDelayParser{},
		RandomParser{},
		IfNotKeyParser{},
		SoundTimerParser{},
		BinaryCodedDecimalParser{},
		LoadRegistersParser{},
		SpriteLocationParser{},
	}

	for _, parser := range parsers {
		if parser.Matches(opCode) {
			return parser.CreateOp(opCode)
		}
	}

	return UnknownOp{}
}