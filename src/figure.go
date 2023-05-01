package main

import (
  "github.com/IKrehan/go_game/util"
)

type Figure interface {
  getBufferSize()
  getFlattenedPoints()
}

type Figure struct {
  vertices [][]float32
  type uint32
}

func getBufferSize(fig Figure) {
  return len(fig.vertices) * 3
}

func getFlattenedPoints(fig Figure) {
  return util.Flatten(fig.vertices)
}
