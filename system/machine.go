package system


type VirtualMachine struct {
	Memory [4096]byte
	Registers [16]byte
	Stack []uint16

	ProgramCounter uint16
	IndexRegister uint16

	DelayTimer byte
	SoundTimer byte

	// Represents the state of key presses
	Keyboard [16]bool
	Pixels [32]int64
}

