package sublist

type Relation string

func Sublist(one []int, two []int) Relation {
	if (len(two) > len(one)){
		return "sublist"
	}
	if(len(one) > len(two)){
		return "superlist"
	}
	return "equal"
}