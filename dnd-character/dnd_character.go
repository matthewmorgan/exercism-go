package dndcharacter

import (
	"math/rand"
	"time"
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
	panic("Please implement the Modifier() function")
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	rand.Seed(time.Now().UnixNano())
	randomInt := 3 + rand.Intn(6) + rand.Intn(6) + rand.Intn(6)
	return randomInt
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	panic("Please implement the GenerateCharacter() function")
}
