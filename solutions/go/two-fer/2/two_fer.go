package twofer

import "fmt"

// Returns a formatted string of who you are sharing with
func ShareWith(name string) string {
  if len(name) == 0 {
    name = "you"
  }
	return fmt.Sprintf("One for %s, one for me.", name)
}
