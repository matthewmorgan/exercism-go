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

type Vigenere struct {
	key string
	currentIndex int
	shiftByIndex []int
}

var INVALID_SHIFT_ARGS = []int{-27, -26, 0, 26, 27}
var re = regexp.MustCompile(`(?m)[a-z]+`)


func NewCaesar() Caesar {
	var caesar Caesar
	caesar.offset = 3
	return caesar
}

func NewShift(offset int) *Caesar {
	for _, invalidArg := range INVALID_SHIFT_ARGS {
		if offset == invalidArg {
			return nil
		}
	}
	return  &Caesar{offset:byte(offset)}
}

func NewVigenere(key string) *Vigenere {
	var v Vigenere
	// Must not be zero length.
	if len(key) == 0 {
		return nil
	}
	// Must not contain non-lowercase alpha characters.
	if clean(key) != key {
		return nil
	}
	var shiftByIndex = make([]int, len(key))
	var checksum = 0
	for idx, letter := range key {
		var offset = int(letter - 'a')
		shiftByIndex[idx] = offset
		checksum += offset
	}
	// Zero checksum indicates a key that is all "a", which is not valid.
	if checksum == 0 {
		return nil
	}
	v.shiftByIndex = shiftByIndex
	v.key = key
	return &v
}

func (v Vigenere) Encode(plain string) string {
	var cleaned = clean(plain)
	var length = len(cleaned)
	var result = make([]byte, length)
	for idx, letter := range []byte(cleaned) {
		result[idx] = xcode(letter, byte(v.shiftByIndex[v.currentIndex]))
		v.incrementIndex()
	}
	return string(result)
}

func (v Vigenere) Decode(encoded string) string {
	var length = len(encoded)
	var result = make([]byte, length)
	for idx, letter := range []byte(encoded) {
		result[idx] = xcode(letter, -byte(v.shiftByIndex[v.currentIndex]))
		v.incrementIndex()
	}
	return string(result)
}

func (v *Vigenere) incrementIndex() {
	if v.currentIndex < len(v.key) - 1 {
		v.currentIndex += 1
	} else {
		v.currentIndex = 0
	}
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

