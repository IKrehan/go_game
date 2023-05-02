package main

import (
  "runtime"
  "github.com/IKrehan/go_game/src/initializer"
  "github.com/go-gl/gl/v4.1-core/gl"
  "github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
  runtime.LockOSThread()

  window := initializer.Glfw(500, 500, "Game of Life")
  defer glfw.Terminate()

  program := initializer.OpenGL()

  figures := []Figure{
    Figure{
      vertices: [][]float32{
        []float32{0, 0.5, 0},     // top
        []float32{-0.5, -0.5, 0}, // left
        []float32{0.5, -0.5, 0},  // right
      }, 
    },
  }
  for !window.ShouldClose() {
    draw(window, program, figures)
  }
}

func draw(window *glfw.Window, program uint32, figs []Figure) {
  gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
  gl.UseProgram(program)


  for _, fig := range figs{
    gl.BindVertexArray(fig.makeVAO())
    gl.DrawArrays(gl.TRIANGLES, 0, len(fig.vertices))
  }

  glfw.PollEvents()
  window.SwapBuffers()
}

