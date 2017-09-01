package system

import (
	"fmt"
)

// Renderer is responsible for rendering pixels and handling user input
type Renderer interface {
	Start(vm *VirtualMachine)
}

// OpCode represents an instruction for the virtual machine
type OpCode uint16

func (o OpCode) String() string {
	return fmt.Sprintf("%04X", uint16(o))
}

// Display represents a 64x32 pixel matrix
type Display [32]uint64

// PixelSetAt determines if the pixel located at coordinate (x, y) is on
func (d *Display) PixelSetAt(x int, y int) bool {
	columnFilter := uint64(1) << (63 - uint(x))
	return d[y]&columnFilter == columnFilter
}

// VirtualMachine the core CHIP8 architecture, containing memory, registers, input, and pixel data
// For reference, see:  http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#2.0
type VirtualMachine struct {
	Memory    [4096]byte // http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#2.1
	Registers [16]byte   // Abbreviated as V0-VF:  http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#2.2
	Stack     []uint16

	ProgramCounter uint16 // Abbreviated as PC
	IndexRegister  uint16 // Abbreviated as I:  http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#2.2

	DelayTimer byte // Abbreviated as DT
	SoundTimer byte // Abbreviated as ST

	// Represents the state of key presses  - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#2.3
	Keyboard [16]bool
	// The state of the pixels - http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#2.4
	Pixels Display

	// Should the machine be running
	Running bool
}

// NewVirtualMachine creates a new virtual machine, loading fonts at the start of the memory space
func NewVirtualMachine() VirtualMachine {
	vm := VirtualMachine{}
	vm.Running = true

	// Load the memory from the font set into the
	// lower "program" memory space
	for i := 0; i < len(fontSet); i++ {
		vm.Memory[i] = fontSet[i]
	}

	return vm
}

// Load data into memory.  By convention, 0x0 - 0x200 is reserved, so data is loaded
// starting at memory address 0x200
func (vm *VirtualMachine) Load(data []byte) {

	// Load the memory starting in application space
	for i := 0; i < len(data); i++ {
		vm.Memory[512+i] = data[i]
	}

	vm.ProgramCounter = 512
}

// OpCodeAt returns an op code at the given memory address
func (vm *VirtualMachine) OpCodeAt(address uint16) OpCode {
	firstByte := uint16(vm.Memory[address])
	secondByte := uint16(vm.Memory[address+1])

	return OpCode((firstByte << 8) + secondByte)
}

// CurrentOpCode returns the op code referenced by the program counter register
func (vm *VirtualMachine) CurrentOpCode() OpCode {
	return vm.OpCodeAt(vm.ProgramCounter)
}

// IncrementPC advances the program counter to reference the next op code
func (vm *VirtualMachine) IncrementPC() {
	vm.ProgramCounter += 2
}

// DecrementTimers decrements delay timer and the sound timer if they are positive
func (vm *VirtualMachine) DecrementTimers() {
	if vm.DelayTimer > 0 {
		vm.DelayTimer--
	}

	if vm.SoundTimer > 0 {
		vm.SoundTimer--
	}
}


// CHIP-8 Font Set.
// See here for reference:  http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#2.4
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
