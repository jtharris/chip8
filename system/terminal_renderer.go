package system

import (
	"github.com/nsf/termbox-go"
	"time"
)

// TerminalRenderer renders a virtual machine to a terminal using termbox
type TerminalRenderer struct {
	shouldQuit bool
}

// Start the render loop, terminating when the escape key is pressed
func (t *TerminalRenderer) Start(vm *VirtualMachine) {
	err := termbox.Init()

	if err != nil {
		panic(err)
	}

	for !t.shouldQuit {
		t.updateKeys(vm)
		t.render(&vm.Pixels)
		time.Sleep(time.Millisecond)
	}

	termbox.Close()
}

func (t *TerminalRenderer) render(d *Display) {
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	drawBorder()

	for col := 0; col < 64; col++ {
		for row := range d {
			if d.PixelSetAt(col, row) {
				termbox.SetCell(col+1, row+1, ' ', termbox.ColorGreen, termbox.ColorGreen)
			}
		}
	}

	termbox.Flush()
}

func drawBorder() {
	for col := 0; col < 66; col++ {
		termbox.SetCell(col, 0, '\u2550', termbox.ColorGreen, termbox.ColorDefault)
		termbox.SetCell(col, 33, '\u2550', termbox.ColorGreen, termbox.ColorDefault)
	}

	for row := 1; row < 33; row++ {
		termbox.SetCell(0, row, '\u2551', termbox.ColorGreen, termbox.ColorDefault)
		termbox.SetCell(65, row, '\u2551', termbox.ColorGreen, termbox.ColorDefault)
	}
}

var termKeyMap = []rune{'x', '1', '2', '3', 'q', 'w', 'e', 'a', 's', 'd', 'z', 'c', '4', 'r', 'f', 'v'}

func (t *TerminalRenderer) updateKeys(vm *VirtualMachine) {
	// Gather up all the keyboard events for 2ms then exit
	time.AfterFunc(time.Millisecond*2, termbox.Interrupt)

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
