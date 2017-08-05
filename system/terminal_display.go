package system

import (
	"github.com/nsf/termbox-go"
	"time"
)

type TerminalDisplay struct {}

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

	for row := range vm.Pixels {
		for col := uint64(0); col < 64; col++  {
			// TODO:  Push this bit logic into a generic display struct?
			columnFilter := uint64(1) << (63 - col)
			if vm.Pixels[row] & columnFilter == columnFilter {
				termbox.SetCell(int(col), row, '\u2588', termbox.ColorGreen, termbox.ColorGreen)
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
	time.AfterFunc(time.Millisecond * 5, termbox.Interrupt)

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