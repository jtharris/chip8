package system

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
)

// OpenGL needs to run on the main OS Thread or bad things happen...
func init() {
	runtime.LockOSThread()
}

type OpenGLDisplay struct {}

func (d *OpenGLDisplay) Start(vm *VirtualMachine) {
	err := glfw.Init()

	if err != nil {
		panic(err)
	}

	// Set up the GLFW window
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

	// Set the coordinate system to be 64x32 "pixels" and set colors
	gl.Ortho(0, 64, 32, 0, 0, 1)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Color3f(0.0, 1.0, 0.0)

	// Main GL loop
	for !window.ShouldClose() {
		d.updateKeys(window, vm)
		d.render(vm)
		window.SwapBuffers()
		glfw.PollEvents()
	}

	glfw.Terminate()
}

// Render the current state of pixels in the virtual machine
func (d *OpenGLDisplay) render(vm *VirtualMachine) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
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

// Hardcoded mapping of qwerty keys to the CHIP8 hex keys - the index of the list (0-15) is the hex key being mapped
// So, for example, hex key 0xA holds the state of the "Z" key
var keyMap = []glfw.Key {glfw.KeyX, glfw.Key1, glfw.Key2, glfw.Key3, glfw.KeyQ, glfw.KeyW, glfw.KeyE, glfw.KeyA,
	glfw.KeyS, glfw.KeyD, glfw.KeyZ, glfw.KeyC, glfw.Key4, glfw.KeyR, glfw.KeyF, glfw.KeyV,
}

// Update the state of the input key map on the virtual machine
func (d *OpenGLDisplay) updateKeys(window *glfw.Window, vm *VirtualMachine) {
	for hex, input := range keyMap {
		vm.Keyboard[hex] = window.GetKey(input) == glfw.Press
	}
}
