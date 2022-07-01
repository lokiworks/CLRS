package main

/**
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

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

func main() {

	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	r := addTwoNumbers(l1, l2)
	println(r)
}
