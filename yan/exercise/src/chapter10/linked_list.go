package main

type Node struct {
	prev *Node
	next *Node
	key  int
}
type LinkedList struct {
	head *Node
	tail *Node
}

func (L *LinkedList) Search(k int) *Node {
	x := L.head
	for x != nil && x.key != k {
		x = x.next
	}
	return x
}

func (L *LinkedList) Prepend(x *Node) {
	x.next = L.head
	x.prev = nil
	if L.head != nil {
		L.head.prev = x
	}
	L.head = x
}

func (L *LinkedList) Insert(x, y *Node) {
	x.next = y.next
	x.prev = y
	if y.next != nil {
		y.next.prev = x
	}
	y.next = x
}

func (L *LinkedList) Delete(x *Node) {
	if x.prev != nil {
		x.prev.next = x.next
	} else {
		L.head = x.next
	}
	if x.next != nil {
		x.next.prev = x.prev
	}
}
