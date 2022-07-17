package main

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	var r *ListNode
	var c *ListNode
	for list1 != nil && list2 != nil {
		t := ListNode{}
		if list1.Val < list2.Val {
			t.Val = list1.Val
			list1 = list1.Next

		} else {
			t.Val = list2.Val
			list2 = list2.Next
		}

		if r == nil {
			r = &t
			c = r
		} else {
			c.Next = &t
			c = c.Next
		}

	}
	if list1 != nil {
		if r == nil {
			r = list1
		} else {
			c.Next = list1
		}
	}

	if list2 != nil {
		if r == nil {
			r = list2
		} else {
			c.Next = list2
		}
	}
	return r
}
