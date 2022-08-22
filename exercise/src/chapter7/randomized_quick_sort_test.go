package main

import (
	"reflect"
	"testing"
)

func Test_randomizedQuickSort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"WithNilArray",
			args{a: nil},
			nil,
		},
		{
			"WithEmptyArray",
			args{a: []int{}},
			[]int{},
		},

		{
			"WithSortedAndDuplicateArray",
			args{a: []int{1, 1, 2, 2}},
			[]int{1, 1, 2, 2},
		},
		{
			"WithUnSortedAndNoDuplicateArray",
			args{a: []int{5, 4, 3, 2, 1}},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"WithUnSortedAndDuplicateArray",
			args{a: []int{5, 4, 2, 2, 1}},
			[]int{1, 2, 2, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randomizedQuickSort(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("randomizedQuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
