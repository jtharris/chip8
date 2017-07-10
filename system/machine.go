package system


type VirtualMachine struct {
	Memory [4096]byte
	Registers [16]byte
	Stack []uint16

	ProgramCounter uint16
	IndexRegister uint16

	DelayTimer byte
	SoundTimer byte

	// TODO:  Figure out what this should be
	Display int
}

