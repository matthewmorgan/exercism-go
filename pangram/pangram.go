package pangram

import "strings"

func IsPangram(phrase string) bool {
	lower := strings.ToLower(phrase)
	foundLetters := buildCleanLetterMap()
	for _, letter := range strings.Split(lower, ""){
		if foundLetters[letter]{
			delete(foundLetters, letter)
			if len(foundLetters) == 0 {
				return true
			}
		}
	}
	return false
}

func buildCleanLetterMap() map[string]bool {
	letters := map[string]bool{}
	for _, letter := range strings.Split("abcdefghijklmnopqrstuvwxyz", ""){
		letters[letter] = true
	}
	return letters
}

