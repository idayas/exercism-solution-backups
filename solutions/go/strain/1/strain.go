package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
type genTypes interface {
  int | float64 | string | ~[]int
}
type genericFunc[c genTypes] func(c) bool

func Keep[T genTypes](inp []T, callback genericFunc[T]) []T {
  tempSlice := []T{}

  for _, item := range inp {
    if callback(item) {
      tempSlice = append(tempSlice, item)
    }
  }

  return tempSlice
}

func Discard[T genTypes](inp []T, callback genericFunc[T]) []T {
  tempSlice := []T{}

  for _, item := range inp {
    if !callback(item) {
      tempSlice = append(tempSlice, item)
    }
  }

  return tempSlice
}
