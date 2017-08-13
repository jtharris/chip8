package system

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

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
		d.UpdateKeys(window, vm)
		d.Render(vm)
		window.SwapBuffers()
		glfw.PollEvents()
	}

	glfw.Terminate()
}

func (d OpenGLDisplay) Render(vm *VirtualMachine) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Color3f(0.0, 1.0, 0.0)
	gl.Begin(gl.QUADS)
	for col := 0; col < 64; col++ {
		for row := range vm.Pixels {
			if vm.PixelSetAt(col, row) {
				c := int32(col)
				r := int32(row)
				gl.Vertex2i(c, r)
				gl.Vertex2i(c+1, r)
				gl.Vertex2i(c+1, r+1)
				gl.Vertex2i(c, r+1)
			}
		}
	}
	gl.End()
}

var keyMap = []glfw.Key {
	glfw.KeyX,
	glfw.Key1,
	glfw.Key2,
	glfw.Key3,
	glfw.KeyQ,
	glfw.KeyW,
	glfw.KeyE,
	glfw.KeyA,
	glfw.KeyS,
	glfw.KeyD,
	glfw.KeyZ,
	glfw.KeyC,
	glfw.Key4,
	glfw.KeyR,
	glfw.KeyF,
	glfw.KeyV,
}

func (d OpenGLDisplay) UpdateKeys(window *glfw.Window, vm *VirtualMachine) {
	for hex, input := range keyMap {
		vm.Keyboard[hex] = window.GetKey(input) == glfw.Press
	}
}
