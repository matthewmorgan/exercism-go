package strain

type Ints []int
type Lists [][]int
type Strings []string

func (collection Ints) Keep(op func (int) bool) Ints {
	if collection == nil {
		return nil
	}
	results := Ints{}
	for _, el := range collection {
		if op(el) {
			results = append(results, el)
		}
	}
	return results
}

func (collection Ints) Discard(op func (int) bool) Ints {
	return collection.Keep(func (el int) bool { return !op(el)})
}

func (collection Lists) Keep(op func ([]int) bool) Lists {
	if collection == nil {
		return nil
	}
	results := Lists{}
	for _, el := range collection {
		if op(el) {
			results = append(results, el)
		}
	}
	return results
}

func (collection Strings) Keep(op func (string) bool) Strings {
	if collection == nil {
		return nil
	}
	results := Strings{}
	for _, el := range collection {
		if op(el) {
			results = append(results, el)
		}
	}
	return results
}


