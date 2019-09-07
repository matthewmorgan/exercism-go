package sublist

type Relation string

func Sublist(one []int, two []int) Relation {
	if equal(one, two){
		return "equal"
	}
	if isSublist(one, two) {
		return "sublist"
	}
	if isSublist(two, one) {
		return "superlist"
	}
	return "unequal"
}

func equal(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for idx, el := range a {
		if el != b[idx] {
			return false
		}
	}
	return true
}

func isSublist(a []int, b []int) bool {
	return len(b) >= len(a) && (equal(a, b[0:len(a)]) || isSublist(a, b[1:]))
}
