package darts

import "math"

func getDistance(x, y float64) float64 {
  return math.Sqrt((x*x) + (y*y))
}

func Score(x, y float64) int {
  d := getDistance(x, y)
  score := 0

  if d <= 1 {
    score = 10
  } else if d <= 5 {
    score = 5
  } else if d <= 10 {
    score = 1
  }

  return score
}
