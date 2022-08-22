package main

import "testing"

func TestQueue_Enqueue(t *testing.T) {
	var queue Queue
	for i := 0; i < 10; i++ {
		queue.Enqueue(i + 1)
	}

	for i := 0; i < 10; i++ {
		element := queue.Dequeue()
		if i+1 != element {
			t.Errorf("actual %q not equal to expected %q", i, element)
		}
	}
}
