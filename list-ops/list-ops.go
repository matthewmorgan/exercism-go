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
	for _, el := range list {
		acc = op(acc, el)
	}
	return acc
}