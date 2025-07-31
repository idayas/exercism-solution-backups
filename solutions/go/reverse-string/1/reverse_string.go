package reverse

import "slices"

func Reverse(input string) string {
  s := []rune(input)
  slices.Reverse(s)
  return string(s)
}
