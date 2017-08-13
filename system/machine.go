package system

import (
	"encoding/hex"
)

// TODO:  Is this the right place for this?
type OpCode uint16
func(o OpCode) String() string {
	bytes := []byte{byte(uint16(o) >> 8), byte(o)}
	return hex.EncodeToString(bytes)
}

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
	// The state of the pixels, rendered to a display
	Pixels [32]uint64
}

func NewVirtualMachine() VirtualMachine {
	// TODO:  Explain memory layout here...
	vm := VirtualMachine{}

	// Load the memory from the font set into the
	// lower "program" memory space
	for i := 0; i < len(fontSet); i++ {
		vm.Memory[i] = fontSet[i]
	}

	return vm
}

func (vm *VirtualMachine) Load(data []byte) {

	// Load the memory starting in application space
	for i := 0; i < len(data); i++ {
		vm.Memory[512 + i] = data[i]
	}

	vm.ProgramCounter = 512
}

func (vm *VirtualMachine) OpCodeAt(address uint16) OpCode {
	firstByte := uint16(vm.Memory[address])
	secondByte := uint16(vm.Memory[address + 1])

	return OpCode((firstByte << 8) + secondByte)
}

func (vm *VirtualMachine) CurrentOpcode() OpCode {
	return vm.OpCodeAt(vm.ProgramCounter)
}

func (vm *VirtualMachine) IncrementPC() {
	vm.ProgramCounter += 2
}

func (vm *VirtualMachine) DecrementTimers() {
	if vm.DelayTimer > 0 {
		vm.DelayTimer--
	}

	if vm.SoundTimer > 0 {
		vm.SoundTimer--
	}
}

func (vm *VirtualMachine) PixelSetAt(x int, y int) bool {
	columnFilter := uint64(1) << (63 - uint(x))
	return vm.Pixels[y] & columnFilter == columnFilter
}


// CHIP-8 Font Set.
var fontSet = [80]byte{
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0xF0, // C
	0xE0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}