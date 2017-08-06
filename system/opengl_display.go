package system

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type OpenGLDisplay struct {}

func (d OpenGLDisplay) Start(vm *VirtualMachine) {
	err := glfw.Init()

	if err != nil {
		panic(err)
	}

	window, err := glfw.CreateWindow(800, 600, "Chip8", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	window.Focus()

	err = gl.Init()
	if err != nil {
		panic(err)
	}

	gl.Ortho(0, 64, 32, 0, 0, 1)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)

	for !window.ShouldClose() {
		UpdateKeys(window, vm)
		Render(vm)
		window.SwapBuffers()
		glfw.PollEvents()
	}

	glfw.Terminate()
}

func Render(vm *VirtualMachine) {
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.Begin(gl.QUADS)
	gl.Color3f(0.0, 1.0, 0.0)
	for col := uint64(0); col < 64; col++  {
		columnFilter := uint64(1) << (63 - col)
		c := int32(col)
		for row := range vm.Pixels {
			if vm.Pixels[row] & columnFilter == columnFilter {
				r := int32(row)
				gl.Vertex2i(c, r)
				gl.Vertex2i(c + 1, r)
				gl.Vertex2i(c + 1, r + 1)
				gl.Vertex2i(c, r + 1)
			}
		}
	}
	gl.End()
}

func UpdateKeys(window *glfw.Window, vm *VirtualMachine) {
	vm.Keyboard[0x1] = window.GetKey(glfw.KeyW) == glfw.Press
	vm.Keyboard[0x4] = window.GetKey(glfw.KeyS) == glfw.Press
	vm.Keyboard[0xC] = window.GetKey(glfw.KeyUp) == glfw.Press
	vm.Keyboard[0xD] = window.GetKey(glfw.KeyDown) == glfw.Press
}
