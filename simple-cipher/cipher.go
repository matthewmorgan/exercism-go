package cipher

import (
	"regexp"
	"strings"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type Caesar struct {}

func NewCaesar() Caesar {
	return Caesar{}
}

var re = regexp.MustCompile(`(?m)[a-z]+`)
const offset = 3

func (c Caesar) Encode(plain string) string{
	var cleaned = clean(plain)
	var length = len(cleaned)
	var result = make([]byte, length)
	for idx, letter := range []byte(cleaned) {
		var shiftedLetter = letter + offset
		if shiftedLetter > 122 {
			shiftedLetter -= 26
		}
		result[idx] = shiftedLetter
	}
	return string(result)
}

func (c Caesar) Decode(encoded string) string{
	var length = len(encoded)
	var result = make([]byte, length)
	for idx, letter := range []byte(encoded) {
		var shiftedLetter = letter - offset
		if shiftedLetter < 97 {
			shiftedLetter += 26
		}
		result[idx] = shiftedLetter
	}
	return string(result)
}

// Clean and normalize input text to lc alpha only
func clean(plain string) string{
	lower := strings.ToLower(plain)
	return strings.Join(re.FindAllString(lower, -1), "")
}

