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

// TODO:  Pull these into a factory struct
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
	IfKeyParser{},
	IfNotRegisterParser{},
}

var ops = make(map[system.OpCode] Operation)

func CreateOperation(opcode system.OpCode) Operation {
	op, ok := ops[opcode]

	if !ok {
		for _, parser := range parsers {
			if parser.Matches(opcode) {
				op = parser.CreateOp(opcode)
				break
			}
		}

		if op == nil {
			op = UnknownOp{code: opcode}
		}

		ops[opcode] = op
	}

	return op
}