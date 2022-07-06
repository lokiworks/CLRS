package chapter1

import (
	"reflect"
	"testing"
)

func Test_insertSort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		// TODO: Add test cases.
		{
			"WithEmptyArray",
			args{a: []int{}},
			[]int{},
		},

		{
			"WithSortedArray",
			args{a: []int{1, 2, 3}},
			[]int{1, 2, 3},
		},

		{
			"WithUnSortedArray",
			args{a: []int{3, 2, 1}},
			[]int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertSort(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
