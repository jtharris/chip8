package system

import (
	"log"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"fmt"
)

func Start() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(800, 600, "Chip8", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	window.Focus()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.Begin(gl.QUADS)
	gl.Color3f(0.0, 1.0, 0.0)
	gl.Vertex2f(-0.8,-0.8)
	gl.Vertex2f(-0.8,0.8)
	gl.Vertex2f(0.8,0.8)
	gl.Vertex2f(0.8,-0.8)
	gl.End()

	window.SwapBuffers()

	for !window.ShouldClose() {
		//gl.Clear(gl.COLOR_BUFFER_BIT)

		// Update
		//window.SwapBuffers()
		glfw.PollEvents()
	}
}
