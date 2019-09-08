package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

var re = regexp.MustCompile(`[^\w|']`)

func WordCount(input string) Frequency {
	counts := Frequency{}
	cleaned := re.ReplaceAllLiteralString(strings.ToLower(input), " ")
	for _, word := range strings.Fields(cleaned) {
		unquoted := strings.Trim(word, `'`)
		counts[unquoted] += 1
	}
	return counts
}