// Package hamming measures hamming distance between two DNA strands
package hamming

import "errors"
import "regexp"

// Distance takes two strings and checks to see if the strings differ at any indicies
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Strands cannot be different lengths!")
	}
	if !allValidChars(a) || !allValidChars(b) {
		return 0, errors.New("Strands cannot contain illegal characters!")
	}
	distance := 0
	for idx := 0; idx < len(a); idx++ {
		if a[idx] != b[idx] {
			distance += 1
		}
	}
	return distance, nil
}

func allValidChars(a string) bool {
	return regexp.MustCompile(`^[ACGT]*$`).MatchString(a)
}
