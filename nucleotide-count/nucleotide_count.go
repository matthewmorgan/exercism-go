// Package dna has a method for counting nucleotides
package dna

import "errors"

type Histogram map[rune]int

type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h :=  Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, nucleotide := range []rune(d) {
		if _, exists := h[nucleotide]; !exists {
			return nil, errors.New("invalid nucleotide")
		}
		h[nucleotide] += 1
	}
	return h, nil
}
