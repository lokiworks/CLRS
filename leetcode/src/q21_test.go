package main

import (
	"reflect"
	"testing"
)

func Test_mergeTwoList(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "WithTwoNilArray",
			args: args{list1: nil, list2: nil},
			want: nil,
		},
		{
			name: "WithList1NilArray",
			args: args{list1: nil, list2: &ListNode{Val: 1, Next: nil}},
			want: &ListNode{Val: 1, Next: nil},
		},

		{
			name: "WithList2NilArray",
			args: args{list1: &ListNode{Val: 1, Next: nil}, list2: nil},
			want: &ListNode{Val: 1, Next: nil},
		},
		{
			name: "WithList1AndList2NotNilArray",
			args: args{list1: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}, list2: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: nil}}}}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: nil}}}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoList(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoList() = %v, want %v", got, tt.want)
			}
		})
	}
}
