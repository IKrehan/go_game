package util

func Flatten[T int32](slice [][]T) []T {
  var flattened []T
  for _, sub_array := range slice {
    for _, item := range sub_array {
      flattened = append(flattened, item)
    }
  }

  return flattened
}
