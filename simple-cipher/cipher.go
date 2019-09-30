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

func doShift(letter byte, offset byte) byte{
	var shiftedLetter = letter + offset
	if shiftedLetter > 122 {
		shiftedLetter -= 26
	} else if shiftedLetter < 97 {
		shiftedLetter +=26
	}
	return shiftedLetter
}

func clean(plain string) string{
	lower := strings.ToLower(plain)
	return strings.Join(re.FindAllString(lower, -1), "")
}

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
	return v.xCode(clean(plain), 1)
}

func (v Vigenere) Decode(encoded string) string {
	return v.xCode(encoded, -1)
}

func (v *Vigenere) incrementIndex() {
	if v.currentIndex < len(v.key) - 1 {
		v.currentIndex += 1
	} else {
		v.currentIndex = 0
	}
}

func (v *Vigenere) xCode(input string, sign int) string {
	var length = len(input)
	var result = make([]byte, length)
	for idx, letter := range []byte(input) {
		result[idx] = doShift(letter, byte(sign * v.shiftByIndex[v.currentIndex]))
		v.incrementIndex()
	}
	return string(result)
}

func (c Caesar) Encode(plain string) string{
	return xCode(clean(plain), c.offset)
}

func (c Caesar) Decode(encoded string) string{
	return xCode(encoded, -c.offset)
}

func xCode(input string, offset byte) string {
	var length = len(input)
	var result = make([]byte, length)
	for idx, letter := range []byte(input) {
		result[idx] = doShift(letter, offset)
	}
	return string(result)
}
