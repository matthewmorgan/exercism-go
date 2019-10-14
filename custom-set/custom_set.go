package stringset

import (
	"strings"
)

// Hash maps are often used under the hood to implement Set in high-level languages.
type Set map[string]bool

func New() *Set {
	return &Set{}
}

func NewFromSlice(input []string) Set {
	var s = Set{}
	// Stuffing these in a map filters for unique values.
	for _, el := range input {
		s.Add(el)
	}

	return s
}

func (s Set) IsEmpty() bool {
	return s.Size() == 0
}

func (s Set) Has(el string) bool {
	_, exists := s[el]
	return exists
}

func (s Set) Size() int {
	return len(s)
}

func Subset(s1 Set, s2 Set) bool {
	if s1.IsEmpty() {
		return true
	}
	if s2.IsEmpty() {
		return false
	}
	if s1.Size() > s2.Size() {
		return false
	}
	for el, _ := range s1 {
		if !s2.Has(el){
			return false
		}
	}
	return true
}

func Disjoint(s1 Set, s2 Set) bool {
	if s1.IsEmpty()|| s2.IsEmpty() {
		return true
	}
	return len(Intersection(s1, s2)) == 0
}

func Equal(s1 Set, s2 Set) bool {
	return s1.Size() == s2.Size() && Intersection(s1, s2).Size() == s1.Size()
}

func (s Set) Add(el string){
	s[el] = true
}

func Union(s1 Set, s2 Set) Set {
	var keys = append(extractKeys(s1), extractKeys(s2)...)
	return NewFromSlice(keys)
}

func Difference(s1 Set, s2 Set) Set {
	if s1.IsEmpty() {
		return Set{}
	}
	var difference = Set{}
	s1, s2 = smallerToLarger(s1, s2)
	for el, _ := range s2 {
		if !s1.Has(el) {
			difference[el] = true
		}
	}
	return difference
}

// Return the set of elements that are in both input sets.
// Iterates over the larger of the two sets, keeping elements that
// are in both sets.
func Intersection(s1 Set, s2 Set) Set {
	if s2.IsEmpty() {
		return Set{}
	}
	var intersect = Set{}
	s1, s2 = smallerToLarger(s1, s2)
	for el, _ := range s2 {
		if s1.Has(el) {
			intersect.Add(el)
		}
		// Exit the loop early if we've already found as many
		// elements in the larger set as there are in the smaller set.
		// This means we can't have further overlap.
		if intersect.Size() == s1.Size() {
			break
		}
	}
	return intersect
}

func (s Set) String() string {
	var result []string
	for _, el := range extractKeys(s) {
		result = append(result, "\"" + el + "\"")
	}
	return "{" + strings.Join(result, ", ") + "}"
}

func smallerToLarger(s1 Set, s2 Set) (Set, Set) {
	if s1.Size() > s2.Size() {
		return s2, s1
	}
	return s1, s2
}

func extractKeys(s Set) []string {
	var keys = make([]string, len(s))
	var index = 0
	for key, _ := range s {
		keys[index] = key
		index += 1
	}
	return keys
}

