package listops

type binFunc func(int, int) int
type unaryFunc func(int) int
type predFunc func(int) bool
type IntList []int

func (list IntList) Foldl(op binFunc, initial int) int {
	acc := initial
	for _, el := range list {
		acc = op(acc, el)
	}
	return acc
}

func (list IntList) Foldr(op binFunc, initial int) int {
	acc := initial
	for i := len(list); i > 0; i-- {
		acc = op(list[i-1], acc)
	}
	return acc
}

func (list IntList) Filter(op predFunc) IntList {
	results := IntList{}
	for _, el := range list {
		if op(el) {
			results = append(results, el)
		}
	}
	return results
}

func (list IntList) Length() int {
	return len(list)
}

func (list IntList) Map(op unaryFunc) IntList {
	results := IntList{}
	for _, el := range list {
		results = append(results, op(el))
	}
	return results
}

func (list IntList) Reverse() IntList {
	listLen := len(list)
	for i := 0; i < listLen/2; i++ {
		temp := list[i]
		swapTarget := listLen - i - 1
		list[i] = list[swapTarget]
		list[swapTarget] = temp
	}
	return list
}

func (list IntList) Append(more IntList) IntList {
	for _, el := range more {
		list = append(list, el)
	}
	return list
}

func (list IntList) Concat(lists []IntList) IntList {
	for _, more := range lists {
		list = list.Append(more)
	}
	return list
}