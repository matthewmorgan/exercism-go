package sublist

import "reflect"

type Relation string

func Sublist(one []int, two []int) Relation {
	if len(two) > len(one) {
		return "sublist"
	}
	if len(one) > len(two) {
		return "superlist"
	}
	if reflect.DeepEqual(one, two) {
		return "equal"
	}
	return "unequal"
}