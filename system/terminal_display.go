package system

import (
	"github.com/nsf/termbox-go"
	"time"
)

type TerminalDisplay struct {
	shouldQuit bool
}

func (t *TerminalDisplay) Start(vm *VirtualMachine) {
	err := termbox.Init()

	if err != nil {
		panic(err)
	}

	for !t.shouldQuit {
		t.UpdateKeys(vm)
		t.Render(vm)
		time.Sleep(time.Millisecond)
	}

	termbox.Close()
}

func (t *TerminalDisplay) Render(vm *VirtualMachine) {
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

	for col := 0; col < 64; col++  {
		for row := range vm.Pixels {
			if vm.PixelSetAt(col, row) {
				termbox.SetCell(col + 1, row + 1, ' ', termbox.ColorGreen, termbox.ColorGreen)
			}
		}
	}

	termbox.Flush()
}


var termKeyMap = []rune {'x', '1', '2', '3', 'q', 'w', 'e', 'a', 's', 'd', 'z', 'c', '4', 'r', 'f', 'v',}

func (t *TerminalDisplay) UpdateKeys(vm *VirtualMachine) {
	// Gather up all the keyboard events for 5ms then exit
	time.AfterFunc(time.Millisecond * 2, termbox.Interrupt)

	for {
		ev := termbox.PollEvent()

		if ev.Type == termbox.EventKey {
			for i, char := range termKeyMap {
				if char == ev.Ch {
					vm.Keyboard[i] = true
					break
				}
			}

			if ev.Key == termbox.KeyEsc {
				t.shouldQuit = true
				return
			}
		} else if ev.Type == termbox.EventInterrupt {
			return
		}
	}
}