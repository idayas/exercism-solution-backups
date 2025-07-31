package dndcharacter

import (
	"math"
	"math/rand"
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
  return int(math.Floor(float64(score - 10) / 2.0))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
  // 15 because its 3 d6, 0 indexed (5 + 5 + 5). Add +1 for each to make it inclusive
  return rand.Intn(15) + 3
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
