// Package scrabble provides a method to score scrabble words
package scrabble

import "strings"

var letterScores = map[string]int{
	"D": 2,
	"G": 2,
	"B": 3,
	"C": 3,
	"M": 3,
	"P": 3,
	"F": 4,
	"H": 4,
	"V": 4,
	"W": 4,
	"Y": 4,
	"K": 5,
	"J": 8,
	"X": 8,
	"Q": 10,
	"Z": 10,
}

// Score takes an input string and returns a score int
func Score(input string) int {
	inputUpper := strings.ToUpper(input)
	score := 0
	for idx := 0; idx < len(inputUpper); idx++ {
		letter := inputUpper[idx : idx+1]
		if val, ok := letterScores[letter]; ok {
			score += val
		} else {
			score += 1
		}
	}
	return score
}
