package main

import (
	"runtime"

	"github.com/IKrehan/go_game/src/figure"
	"github.com/IKrehan/go_game/src/initializer"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
	runtime.LockOSThread()

	window := initializer.Glfw(500, 500, "Game of Life")
	defer glfw.Terminate()

	program := initializer.OpenGL()

	figures := []figure.Figure{
		figure.New(
			[][]float32{
				{0, 0.5, 0},     // top
				{-0.5, -0.5, 0}, // left
				{0.5, -0.5, 0},  // right
			},
		),
	}
	for !window.ShouldClose() {
		draw(window, program, figures)
	}
}

func draw(window *glfw.Window, program uint32, figs []figure.Figure) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	for _, fig := range figs {
		gl.BindVertexArray(fig.MakeVAO())
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(fig.Vertices)))
	}

	glfw.PollEvents()
	window.SwapBuffers()
}
