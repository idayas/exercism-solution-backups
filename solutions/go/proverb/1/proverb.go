// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Writes a proverb for an array of variable length
func Proverb(rhyme []string) []string {
	proverb := []string{}
	if len(rhyme) == 0 {
		return proverb
	}
	finalLine := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	
	for i := 0; i < (len(rhyme) - 1); i++ {
		newLine := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i + 1])
		proverb = append(proverb, newLine)
	}
	proverb = append(proverb, finalLine)
	return proverb
}
