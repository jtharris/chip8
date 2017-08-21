package operations

import (
	"chip8/system"
)

type operationParser interface {
	matches(opcode system.OpCode) bool
	createOp(opcode system.OpCode) Operation
}

type Operation interface {
	String() string
	Execute(machine *system.VirtualMachine)
}

var parsers = []operationParser{
	clearParser{},
	returnParser{},
	gotoParser{},
	callParser{},
	ifConstantParser{},
	ifNotConstantParser{},
	ifRegisterParser{},
	assignConstantParser{},
	addConstantParser{},
	assignRegisterParser{},
	bitwiseOrParser{},
	bitwiseAndParser{},
	bitwiseXorParser{},
	addRegisterParser{},
	subtractRegisterParser{},
	shiftRightParser{},
	shiftLeftParser{},
	reverseSubtractRegisterParser{},
	setIndexParser{},
	drawParser{},
	delayTimerParser{},
	setToDelayParser{},
	randomParser{},
	ifNotKeyParser{},
	soundTimerParser{},
	binaryCodedDecimalParser{},
	loadRegistersParser{},
	spriteLocationParser{},
	addToIndexParser{},
	ifKeyParser{},
	ifNotRegisterParser{},
	dumpRegistersParser{},
	getKeyParser{},
}

var ops = make(map[system.OpCode] Operation)

func CreateOperation(opcode system.OpCode) Operation {
	op, ok := ops[opcode]

	if !ok {
		for _, parser := range parsers {
			if parser.matches(opcode) {
				op = parser.createOp(opcode)
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