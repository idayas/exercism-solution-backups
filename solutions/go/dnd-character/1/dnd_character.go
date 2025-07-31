package dndcharacter

import (
	"math"
	"math/rand"
	"sort"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
  s := float64(score)
  return int(math.Floor((s - 10.0) / 2.0))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
  rolls := []int{0,0,0,0}
  for i := range rolls {
    rolls[i] = rand.Intn(6) + 1
  }
  sort.Ints(rolls)
  score := 0;

  for i, v := range rolls {
    if i == 0 {continue}
    score += v
  }

  return score
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
  con := Ability();
  return Character{
    Strength: Ability(),
    Dexterity: Ability(),
    Constitution: con,
    Intelligence: Ability(),
    Wisdom: Ability(),
    Charisma: Ability(),
    Hitpoints:  10 + Modifier(con),
  }
}
