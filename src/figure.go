package main

import (
  "github.com/IKrehan/go_game/src/util"
)

type Figure interface {
  getBufferSize()
  getFlattenedPoints()
  makeVAO()
}

type Figure struct {
  vertices [][]float32
}

func getBufferSize(fig Figure) {
  return len(fig.getFlattenedPoints()) * 4
}

func getFlattenedPoints(fig Figure) {
  return util.Flatten(fig.vertices)
}

func makeVAO(fig Figure) uint32 {
  var vbo uint32
  gl.GenBuffers(1, &vbo)
  gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
  gl.BufferData(gl.ARRAY_BUFFER, fig.getBufferSize(), gl.Ptr(fig.getFlattenedPoints()), gl.STATIC_DRAW)

  var vao uint32
  gl.GenVertexArrays(1, &vao)
  gl.BindVertexArray(vao)
  gl.EnableVertexAttribArray(0)
  gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
  gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

  return vao
}

