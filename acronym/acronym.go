// Package acronym converts a phrase to an acronym
package acronym

import (
	"regexp"
	"strings"
)

// take a string s and return an acronym
func Abbreviate(s string) string {
	acronym := ""
	for _, token := range tokenize(s) {
		acronym += strings.ToUpper(string(token[0]))
	}
	return acronym
}

func tokenize(phrase string) []string {
	nocommas := regexp.MustCompile(`[,]`).ReplaceAllString(phrase, "")
	nodashes := regexp.MustCompile(`[-]`).ReplaceAllString(nocommas, " ")
	return strings.Split(nodashes, " ")
}
