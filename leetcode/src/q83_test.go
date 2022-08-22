package main

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			"With112",
			args{&ListNode{1, &ListNode{1, &ListNode{2, nil}}}},
			&ListNode{1, &ListNode{2, nil}},
		},
		{
			"With11233",
			args{&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}},
			&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteDuplicatesV2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			"With112",
			args{&ListNode{1, &ListNode{1, &ListNode{2, nil}}}},
			&ListNode{1, &ListNode{2, nil}},
		},
		{
			"With11233",
			args{&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}},
			&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicatesV2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicatesV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
