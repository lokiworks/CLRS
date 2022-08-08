package main

import "testing"

func TestStack_Push(t *testing.T) {
	var stack Stack
	for i := 0; i < 10; i++ {
		stack.Push(i + 1)
	}

	for i := 10; i > 0; i-- {
		element, err := stack.Pop()
		if err != nil {
			if i != element {
				t.Errorf("actual %q not equal to expected %q", i, element)
			}
		}

	}
}
