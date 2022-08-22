package main

import "testing"

func TestLinkedList_Insert(t *testing.T) {
	var list LinkedList
	y := &Node{key: 0}
	list.Prepend(y)
	for i := 1; i < 10; i++ {
		node := Node{key: i + 1}
		list.Insert(&node, y)
	}

	for i := 0; i <= 10; i++ {
		node := list.Search(i)
		if node != nil && node.key != i {
			t.Errorf("actual %q not equal to expected %q", i, node.key)
		}

	}

	for i := 0; i <= 10; i++ {
		list.Delete(&Node{key: i})
	}
	if list.head != nil {
		t.Errorf("expected empty list")
	}
}
