package main

func deleteDuplicates(head *ListNode) *ListNode {
	var r *ListNode
	var c *ListNode
	m := make(map[int]bool)

	for head != nil {
		v := head.Val
		_, ok := m[v]
		if !ok {
			m[v] = true
			if r == nil {
				r = head
				c = r
			} else {
				c.Next = head
				c = c.Next
			}
		}
		head = head.Next
	}
	if c != nil {
		c.Next = nil
	}

	return r

}

// 参考官方实现
func deleteDuplicatesV2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 只使用一个变量进行控制
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
