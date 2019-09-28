package protein

import "errors"

var ErrStop = errors.New("stop sequence encountered")
var ErrInvalidBase = errors.New("invalid base encountered")


// Hashmap lookup is ~2/3 the time of switch-case
var translate = map[string]string{
	"AUG": "Methionine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "Stop",
	"UAG": "Stop",
	"UGA": "Stop",
}

func FromCodon(codon string) (string, error) {
	if protein, exists := translate[codon]; !exists {
		return "", ErrInvalidBase
	} else if protein == "Stop" {
		return "", ErrStop
	} else {
		return protein, nil
	}
}

func FromRNA(rna string) ([]string, error) {
	proteins := []string{}
	for i := 0; i < len(rna)-2; i += 3 {
		protein, err := FromCodon(rna[i : i+3])
		if err != nil {
			if err == ErrStop {
				return proteins, nil
			}
			return proteins, err
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}
