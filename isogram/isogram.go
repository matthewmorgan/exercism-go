// Package isogram checks if words are isograms
package isogram

import "strings"
import "unicode"

// IsIsogram accepts a word as a string and returns a boolean
func IsIsogram(input string) bool {
	set := make(map[rune]bool)
	for _, c := range strings.ToUpper(input) {
		if unicode.IsLetter(c) {
			if _, exists := set[c]; exists {
				return false
			}
			set[c] = true
		}
	}
	return true
}
