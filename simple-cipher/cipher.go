package cipher


type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type Caesar struct {}

func NewCaesar() Caesar {
	return Caesar{}
}

func (c Caesar) Encode(plain string) string{
	var length = len(plain)
	const offset = 3
	var result = make([]byte, length)
	for idx, letter := range []byte(plain) {
		var shiftedLetter = letter + offset
		if shiftedLetter > 122 {
			shiftedLetter -= 26
		}
		result[idx] = shiftedLetter
	}
	return string(result)
}

func (c Caesar) Decode(plain string) string{
	return plain
}

