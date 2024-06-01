package dndcharacter

import (
	"math/rand"
	"sort"
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
// We roll three six-sided dice and sum the top three results
func Ability() int {
	rand.Seed(time.Now().UnixNano())
	randomInts := []int{}
	for i := 0; i < 4; i++ {
		randomInts = append(randomInts, rand.Intn(6)+1)
	}
	sort.Ints(randomInts)
	randomInts = randomInts[1:]
	ability := 0
	for _, i := range randomInts {
		ability += i
	}

	return ability
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	panic("Please implement the GenerateCharacter() function")
}
