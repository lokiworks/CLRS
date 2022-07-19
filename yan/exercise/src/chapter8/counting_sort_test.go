package main

import (
	"reflect"
	"testing"
)

func Test_countingSort(t *testing.T) {
	type args struct {
		a []int
		k int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "WithUnsortedArray",
			args: args{a: []int{2, 5, 3, 0, 2, 3, 0, 3}, k: 5},
			want: []int{0, 0, 2, 2, 3, 3, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countingSort(tt.args.a, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countingSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
