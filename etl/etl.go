// Package provides an exported method to transform legacy data to a new format
package etl

import "strings"

// Transform converts legacy data to new format
func Transform(input map[int][]string) map[string]int {
	results := map[string]int{}

	for score, letters := range input {
		for _, letter := range letters {
			results[strings.ToLower(letter)] += score
		}
	}

	return results
}
