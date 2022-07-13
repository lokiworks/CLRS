package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var l *ListNode
	var n *ListNode

	r := 0
	for l1 != nil && l2 != nil {
		t := l1.Val + l2.Val + r
		r = t / 10
		v := 0

		if r > 0 {

			v = t % 10

		} else {
			v = t

		}

		if l == nil {
			l = &ListNode{v, nil}
			n = l
		} else {
			n.Next = &ListNode{v, nil}
			n = n.Next
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		t := l1.Val + r
		r = t / 10
		v := 0
		if r > 0 {

			v = t % 10

		} else {
			v = t

		}

		if l == nil {
			l = &ListNode{v, nil}
			n = l
		} else {

			n.Next = &ListNode{v, nil}
			n = n.Next
		}

		l1 = l1.Next
	}

	for l2 != nil {
		t := l2.Val + r
		r = t / 10
		v := 0
		if r > 0 {

			v = t % 10

		} else {
			v = t

		}

		if l == nil {
			l = &ListNode{v, nil}
			n = l
		} else {

			n.Next = &ListNode{v, nil}
			n = n.Next
		}

		l2 = l2.Next
	}
	// r
	if r > 0 {
		n.Next = &ListNode{r, nil}
	}

	return l

}
