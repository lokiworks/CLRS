package main

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"WithMatch3",
			args{nums: []int{1, 2, 3}, target: 3},
			[]int{0, 1},
		},
		{
			"WithMatch4",
			args{nums: []int{1, 2, 3}, target: 4},
			[]int{0, 2},
		},
		{
			"WithMismatch3",
			args{nums: []int{1, 3, 4}, target: 3},
			nil,
		},
		{
			"WithMismatch4",
			args{nums: []int{1, 2, 4}, target: 4},
			nil,
		},
		{
			"WithMatch6",
			args{nums: []int{3, 2, 4}, target: 6},
			[]int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
