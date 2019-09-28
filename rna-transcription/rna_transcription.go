package strand

var translate = map[rune]string{
	'C': "G",
	'G': "C",
	'T': "A",
	'A': "U",
}

func ToRNA(dna string) string {
	result := ""
	for _, a := range []rune(dna){
		result += translate[a]
	}
	return result
}
