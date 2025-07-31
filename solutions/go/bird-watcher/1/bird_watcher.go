package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
  // Trying to only use for loop, not for with range
  finalTotal := 0
  for i := 0; i < len(birdsPerDay); i++ {
    finalTotal += birdsPerDay[i]
  }
  return finalTotal
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
  weekStart := (week - 1) * 7
  weekEnd := weekStart + 7
  weekData := birdsPerDay[weekStart:weekEnd]

  weekTotal := 0
  for i := 0; i < len(weekData); i++ {
    weekTotal += weekData[i]
  }

  return weekTotal
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
  
  for i := 0; i < len(birdsPerDay); i++ {
    if i % 2 == 0 {
      birdsPerDay[i] += 1
    }
  }

  return birdsPerDay
}
