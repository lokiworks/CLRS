package main

import "errors"

type Stack []int

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return 0, errors.New("underflow")
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, nil

}
