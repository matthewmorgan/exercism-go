// Package hamming measures hamming distance between two DNA strands
package hamming

import "errors"

// Distance takes two strings and checks to see if the strings differ at any indicies
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Strands cannot be different lengths!")
	}
	distance := 0
	for idx := range a {
		if a[idx] != b[idx] {
			distance += 1
		}
	}
	return distance, nil
}
