// Package acronym converts a phrase to an acronym.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate takes a string s and return an acronym.
func Abbreviate(s string) string {
	var acronym []byte
	for _, token := range tokenize(strings.ToUpper(s)) {
		acronym = append(acronym, token[0])
	}
	return string(acronym)
}

func tokenize(phrase string) []string {
	return regexp.MustCompile("[^A-Z]+").Split(phrase, -1)
}
