package scrabble

import "unicode"

func Score(word string) int {
  fs := 0
  for _, c := range word {
    switch unicode.ToUpper(c) {
    case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
      fs += 1
    case 'D', 'G':
      fs += 2
    case 'B', 'C', 'M', 'P':
      fs += 3
    case 'F', 'H', 'V', 'W', 'Y':
      fs += 4
    case 'K':
      fs += 5
    case 'J', 'X':
      fs += 8
    case 'Q', 'Z':
      fs += 10
    default:
      fs += 0
    }
  }
  return fs
}
