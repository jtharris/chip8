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

var parsers = []OperationParser{
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
	AddToIndexParser{},
}

func CreateOperation(opCode system.OpCode) Operation {
	for _, parser := range parsers {
		if parser.Matches(opCode) {
			return parser.CreateOp(opCode)
		}
	}

	return UnknownOp{code: opCode}
}