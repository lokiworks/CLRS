package main

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			"WithL1IsNilAndL2IsNil",
			args{l1: nil, l2: nil},
			nil,
		},
		{
			"WithL1IsNilAndL2IsZero",
			args{l1: nil, l2: &ListNode{Val: 0, Next: nil}},
			&ListNode{Val: 0, Next: nil},
		},

		{
			"WithL2IsNilAndL1IsZero",
			args{l2: nil, l1: &ListNode{Val: 0, Next: nil}},
			&ListNode{Val: 0, Next: nil},
		},

		{
			"WithL2IsNotNilAndL1IsNotNil",
			args{l1: &ListNode{Val: 1, Next: nil}, l2: &ListNode{Val: 0, Next: nil}},
			&ListNode{Val: 1, Next: nil},
		},

		{
			"WithValGreaterThan10",
			args{l1: &ListNode{Val: 4, Next: nil}, l2: &ListNode{Val: 6, Next: nil}},
			&ListNode{Val: 0, Next: &ListNode{1, nil}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
