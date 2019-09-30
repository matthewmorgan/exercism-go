package cipher

import (
	"regexp"
	"strings"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type Caesar struct {
	offset byte
}

var INVALID_ARGS = []int{-27, -26, 0, 26, 27}
var re = regexp.MustCompile(`(?m)[a-z]+`)


func NewCaesar() Caesar {
	var caesar Caesar
	caesar.offset = 3
	return caesar
}

func NewShift(offset int) *Caesar {
	for _, invalidArg := range INVALID_ARGS {
		if offset == invalidArg {
			return nil
		}
	}
	return  &Caesar{offset:byte(offset)}
}

func (c Caesar) Encode(plain string) string{
	var cleaned = clean(plain)
	var length = len(cleaned)
	var result = make([]byte, length)
	for idx, letter := range []byte(cleaned) {
		result[idx] = xcode(letter, c.offset)
	}
	return string(result)
}

func (c Caesar) Decode(encoded string) string{
	var length = len(encoded)
	var result = make([]byte, length)
	for idx, letter := range []byte(encoded) {
		result[idx] = xcode(letter, -c.offset)
	}
	return string(result)
}

func xcode(letter byte, offset byte) byte{
	var shiftedLetter = letter + offset
	if shiftedLetter > 122 {
		shiftedLetter -= 26
	} else if shiftedLetter < 97 {
		shiftedLetter +=26
	}
	return shiftedLetter
}

// Clean and normalize input text to lc alpha only
func clean(plain string) string{
	lower := strings.ToLower(plain)
	return strings.Join(re.FindAllString(lower, -1), "")
}

