package main

import (
  "log"
  "runtime"
  "strings"
  "fmt"

  "github.com/IKrehan/go_game/util"
  "github.com/IKrehan/go_game/initializer"
  "github.com/go-gl/gl/v4.1-core/gl"
  "github.com/go-gl/glfw/v3.2/glfw"
)

const (
  width  = 500
  height = 500
)

var (
  triangle = []float32{
    0, 0.5, 0,     // top
    -0.5, -0.5, 0, // left
    0.5, -0.5, 0,  // right
  }
)

func main() {
  slice = [][]int32{
    [1, 2, 3],
    [4, 5, 6],
  }

  log.Println(utils.Flatten(slice))
}

func amain() {
  runtime.LockOSThread()

  window := initializers.Glfw()
  defer glfw.Terminate()

  program := initializers.OpenGL()

  vao := makeVAO(triangle)
  for !window.ShouldClose() {
    draw(vao, window, program)
  }
}

func draw(window *glfw.Window, program uint32, vaos []uint32) {
  gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
  gl.UseProgram(program)

  gl.BindVertexArray(vao)
  gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle) / 3))

  glfw.PollEvents()
  window.SwapBuffers()
}

func makeVAO(points []float32) uint32 {
  var vbo uint32
  gl.GenBuffers(1, &vbo)
  gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
  gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

  var vao uint32
  gl.GenVertexArrays(1, &vao)
  gl.BindVertexArray(vao)
  gl.EnableVertexAttribArray(0)
  gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
  gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

  return vao
}

