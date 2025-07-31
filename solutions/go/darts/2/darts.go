package darts

import "math"

func getDistance(x, y float64) float64 {
  return math.Sqrt((x*x) + (y*y))
}

func Score(x, y float64) int {
  d := getDistance(x, y)
  
  switch {
  case d <= 1:
    return 10
  case d <= 5:
    return 5
  case d <= 10:
    return 1
  default:
    return 0
  }

}
