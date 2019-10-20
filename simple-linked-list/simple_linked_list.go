package linkedlist

import "errors"

type Element struct {
	data int
	next *Element
}

type List struct {
	head *Element
	size int
}

func New(input []int) *List {
	var list = List{nil, 0}
	for _, value := range input {
		list.Push(value)
	}
	return &list
}

func (list *List) Size() int {
	return list.size
}

func (list *List) Push(value int) {
	var element = Element{value, list.head}
	element.next = list.head
	list.head = &element
	list.size += 1
}

func (list *List) Pop() (int, error) {
	var element = list.head
	if element == nil {
		return -1, errors.New("empty list")
	}
	list.head = element.next
	return element.data, nil
}

func (list *List) Array() []int {
	var array = []int{}
	var element = list.head
	for {
		if element == nil {
			return array
		}
		array = append([]int{element.data}, array...)
		element = element.next
	}
}
func (list *List) Reverse() *List {
	if list.head == nil {
		return list
	}
	return reverse(list, nil)
}

func reverse(list *List, prev *Element) *List {
	if  list.head.next != nil {
		var current = list.head
		list.head = list.head.next
		current.next = prev
		return reverse(list, current)
	}
	list.head.next = prev
	return list
}
