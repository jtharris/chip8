package system


type VirtualMachine struct {
	Memory [4096]byte
	Registers [16]byte
	Stack [64]byte

	ProgramCounter int
	IndexRegister int16

	// TODO:  Figure out what this should be
	Display int
}

