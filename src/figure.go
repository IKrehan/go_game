package main

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
}


func flatten[T](slice [][]T) T {
  flattened []T = []
  for index, sub_array := range slice {
    for index1, item := range sub_array {
      append(flattened, item)
    }
  }
}
