package util

func Flatten[T](slice [][]T) T {
  flattened []T = []
  for index, sub_array := range slice {
    for index1, item := range sub_array {
      append(flattened, item)
    }
  }
}
