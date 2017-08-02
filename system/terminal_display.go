package system

import "github.com/nsf/termbox-go"

type TerminalDisplay struct{}

func (t TerminalDisplay) Initialize(vm *VirtualMachine) {
	err := termbox.Init()

	if err != nil {
		panic(err)
	}
}

func (t TerminalDisplay) Close() {
	termbox.Close()
}

func (t TerminalDisplay) Render(vm *VirtualMachine) {
	// TODO:  Configure color palette
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

	for row := range vm.Pixels {
		for col := uint64(0); col < 64; col++  {
			// TODO:  Push this bit logic into a generic display struct?
			if vm.Pixels[row] & (col ^ 2) == col ^ 2 {
				termbox.SetCell(int(col), row, '#', termbox.ColorGreen, termbox.ColorBlack)
			}
		}
	}

	termbox.Flush()
}
