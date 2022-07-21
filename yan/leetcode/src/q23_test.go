package main

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {

	li := []*ListNode{}
	l1 := ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}
	l2 := ListNode{Val: 5, Next: nil}
	li = append(li, &l1)
	li = append(li, &l2)
	r := mergeKLists(li)
	e := &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}

	if !reflect.DeepEqual(r, e) {
		t.Errorf("mergeKLists() = %v, want %v", r, e)
	}

}
