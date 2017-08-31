package operations

import (
	"chip8/system"
)

// operationParser indicates if it is able to parse a given opcode, and if so, can translate it to an Operation
type operationParser interface {
	matches(opcode system.OpCode) bool
	createOp(opcode system.OpCode) Operation
}

// Operation encapsulates an instruction for a VirtualMachine
type Operation interface {
	String() string
	Execute(machine *system.VirtualMachine)
}

// Available operationParsers
var parsers = []operationParser{
	addConstantParser{},
	addRegisterParser{},
	addToIndexParser{},
	assignConstantParser{},
	assignRegisterParser{},
	binaryCodedDecimalParser{},
	bitwiseAndParser{},
	bitwiseOrParser{},
	bitwiseXorParser{},
	callParser{},
	clearParser{},
	delayTimerParser{},
	drawParser{},
	dumpRegistersParser{},
	getKeyParser{},
	gotoParser{},
	ifConstantParser{},
	ifKeyParser{},
	ifNotConstantParser{},
	ifNotKeyParser{},
	ifNotRegisterParser{},
	ifRegisterParser{},
	loadRegistersParser{},
	randomParser{},
	returnParser{},
	reverseSubtractRegisterParser{},
	setIndexParser{},
	setToDelayParser{},
	shiftLeftParser{},
	shiftRightParser{},
	soundTimerParser{},
	spriteLocationParser{},
	subtractRegisterParser{},
}

// Cache for OpCode -> Operation mapping, lazily evaluated
var ops = make(map[system.OpCode]Operation)

// CreateOperation is a factory to create the correct Operation for the given OpCode
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
