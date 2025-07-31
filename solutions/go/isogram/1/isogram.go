package isogram

import "strings"

func IsIsogram(word string) bool {
  m := map[rune]bool{}
  for _, c := range strings.ToLower(word) {
    if m[c] && c != ' ' && c != '-' {
      return false
    }
    m[c] = true
  }
  return true
}
