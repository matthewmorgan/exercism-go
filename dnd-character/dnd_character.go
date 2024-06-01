package dndcharacter

import (
	"math"
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
// Subtract 10 from the constituition score, divide by two, and round down.
func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2))
}

// Ability uses randomness to generate the score for an ability
// We roll three six-sided dice and sum the top three results
func Ability() int {
	sum, min := 0, 20
	for i := 0; i < 4; i++ {
		randomInt := rand.Intn(6) + 1
		sum += randomInt
		if randomInt < min {
			min = randomInt
		}
	}
	return sum - min
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	rand.Seed(time.Now().UnixNano())
	character := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}
	character.Hitpoints = 10 + Modifier(character.Constitution)
	return character
}
