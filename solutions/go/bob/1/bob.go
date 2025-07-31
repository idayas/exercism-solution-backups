// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "regexp"

// Hey should have a comment documenting it.
func Hey(remark string) string {
  // Yell Question -> Calm down, I know what I'm doing!
  m, _ := regexp.MatchString(`^[^a-z]*[A-Z]{2,}[^a-z]*[?]$`, remark); if m {
    return "Calm down, I know what I'm doing!"
  }

  // Yell -> Woah, chill out!
  m, _ = regexp.MatchString(`^[^a-z]*[A-Z]{2,}[^a-z]*$`, remark); if m {
    return "Whoa, chill out!"
  }

  // Question -> Sure.

  m, _ = regexp.MatchString(`\?[\t\s\n]*$`, remark); if m {
    return "Sure."
  }

  // Silence -> Fine. Be that way!
  m, _ = regexp.MatchString(`^[\t\n\s]*$`, remark); if m {
    return "Fine. Be that way!"
  }
  // Default -> Whatever.
  return "Whatever."
}
