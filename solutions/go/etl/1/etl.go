package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
  scoring := map[string]int{}
  for k, v := range in {
    for _, s := range v {
      scoring[strings.ToLower(s)] = k
    }
  }
  return scoring;
}
