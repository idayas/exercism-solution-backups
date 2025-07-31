package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

func reverse(str string) (result string) {
  for _, v := range str {
    result = string(v) + result
  }
  return
}


func Valid(id string) bool {
  checkNumbers := regexp.MustCompile(`^[0-9]*$`)
  str := strings.ReplaceAll(id, " ", "")

  if len(str) <= 1 { return false }
  if !checkNumbers.MatchString(str) { return false }
  rev := reverse(str)
  total := 0;

  for i, v := range rev {
    // Convert to int
    num, _ := strconv.Atoi(string(v))

    // We want every other, and Go iterates from 0
    // So 0 == true, 1 == false, 2 == true, ...
    if i % 2 == 0 { 
      total += num
      continue 
    }
    if num * 2 > 9 {
      total += (num * 2) - 9
      continue;
    }
    total += num * 2
  }
  return total % 10 == 0
}
