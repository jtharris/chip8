package system

import (
	"github.com/nsf/termbox-go"
	"time"
)

type TerminalDisplay struct {
	PreviousFrame [32]uint64
}

func (t TerminalDisplay) Initialize() {
	err := termbox.Init()

	if err != nil {
		panic(err)
	}
}

func (t TerminalDisplay) Close() {
	termbox.Close()
}

func (t TerminalDisplay) Render(vm *VirtualMachine) {
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

	// Draw the border
	for col := 0; col < 66; col++ {
		termbox.SetCell(col, 0, '\u2550', termbox.ColorGreen, termbox.ColorDefault)
		termbox.SetCell(col, 33, '\u2550', termbox.ColorGreen, termbox.ColorDefault)
	}

	for row := 1; row < 33; row++ {
		termbox.SetCell(0, row, '\u2551', termbox.ColorGreen, termbox.ColorDefault)
		termbox.SetCell(65, row, '\u2551', termbox.ColorGreen, termbox.ColorDefault)
	}

	// TODO:  Push this bit logic into a generic display struct?
	for col := uint64(0); col < 64; col++  {
		columnFilter := uint64(1) << (63 - col)
		for row := range vm.Pixels {
			if vm.Pixels[row] & columnFilter == columnFilter {
				termbox.SetCell(int(col) + 1, row + 1, ' ', termbox.ColorGreen, termbox.ColorGreen)
			}
		}
	}

	termbox.Flush()
}

func (t TerminalDisplay) UpdateKeys(vm *VirtualMachine) {
	// Clear the keyboard - TODO:  Pull this out somewhere else?
	for i := 0; i < len(vm.Keyboard); i++ {
		vm.Keyboard[i] = false
	}

	// Gather up all the keyboard events for 5ms then exit
	time.AfterFunc(time.Millisecond * 2, termbox.Interrupt)

	for {
		ev := termbox.PollEvent()

		// TODO:  Temp
		if ev.Type == termbox.EventKey {
			switch ev.Ch {
			case 'w':
				vm.Keyboard[0x1] = true
			case 's':
				vm.Keyboard[0x4] = true
			}

			switch ev.Key {
			case termbox.KeyArrowUp:
				vm.Keyboard[0xC] = true
			case termbox.KeyArrowDown:
				vm.Keyboard[0xD] = true
			}
		} else if ev.Type == termbox.EventInterrupt {
			return
		}
	}
}