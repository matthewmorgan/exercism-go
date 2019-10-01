package anagram

import (
	"sort"
	"strings"
)

type maybeAnagram struct {
	lowercase string
	sorted    string
}

func Detect(subject string, candidates []string) []string {
	var parsedSubject = stringToSortedString(subject)
	var anagrams = []string{}

	for _, candidate := range candidates {
		var parsedCandidate = stringToSortedString(candidate)
		if parsedCandidate.isAnagramOf(&parsedSubject) {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}

func (m1 *maybeAnagram) isAnagramOf(m2 *maybeAnagram) bool {
	return m1.sorted == m2.sorted && m1.lowercase != m2.lowercase
}

func stringToSortedString(input string) maybeAnagram {
	var lower = strings.ToLower(input)
	var lettersArray = strings.Split(lower, "")
	sort.Strings(lettersArray)
	var sortedInput = strings.Join(lettersArray, "")
	return maybeAnagram{lower, sortedInput}
}
