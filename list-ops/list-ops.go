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
	for i := len(list); i > 0 ; i-- {
		acc = op(list[i-1], acc)
	}
	return acc
}

func (list IntList) Filter(op predFunc) IntList {
	results := IntList{}
	for _, el := range list {
		if op(el){
			results = append(results, el)
		}
	}
	return results
}