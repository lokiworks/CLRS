package main

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < len(lists); i++ {

		head := lists[i]
		for head != nil {
			heap.Push(h, head.Val)
			head = head.Next
		}

	}

	var l *ListNode
	var c *ListNode

	for h.Len() != 0 {
		v := heap.Pop(h)
		if l == nil {
			l = &ListNode{Val: v.(int), Next: nil}
			c = l
		} else {
			c.Next = &ListNode{Val: v.(int), Next: nil}
			c = c.Next
		}
	}

	return l
}

func main() {
	li := []*ListNode{}
	l1 := ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}
	l2 := ListNode{Val: 5, Next: nil}
	li = append(li, &l1)
	li = append(li, &l2)
	l := mergeKLists(li)
	println(l)

}
