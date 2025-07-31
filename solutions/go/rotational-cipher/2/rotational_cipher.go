package rotationalcipher

import (
	"strings"
	"unicode"
)

var chars = []string{
	"a", "b", "c", "d", "e",
	"f", "g", "h", "i", "j",
	"k", "l", "m", "n", "o",
	"p", "q", "r", "s", "t",
	"u", "v", "w", "x", "y",
	"z",
}

func findCharacter(char string) int {
	loc := -1
	for i, c := range chars {
		if c == char {
			loc = i
		}
	}
	return loc
}

func RotationalCipher(plain string, shiftKey int) string {

	shift := (shiftKey % 26)
	newChars := chars[shift:26]
	newChars = append(newChars, chars[:shift]...)

	var newStr string

	for _, c := range plain {
		str := strings.ToLower(string(c))
		newChar := findCharacter(str)
		newLoc := (newChar + shiftKey) % 26

		if newChar > -1 {
			str = chars[newLoc]
			if unicode.IsUpper(c) {
				str = strings.ToUpper(str)
			}
		}
		newStr += str
	}

	return newStr
}
