package main

import (
	"reflect"
	"testing"
)

func Test_radixSort(t *testing.T) {
	type args struct {
		a []int
		d int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "WithUnsortedArray",
			args: args{a: []int{2, 5, 3, 0, 2, 3, 0, 3}, d: 1},
			want: []int{0, 0, 2, 2, 3, 3, 3, 5},
		},

		{
			name: "WithUnsortedArray2",
			args: args{a: []int{170, 45, 75, 90, 802, 24, 2, 66}, d: 3},
			want: []int{2, 24, 45, 66, 75, 90, 170, 802},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := radixSort(tt.args.a, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("radixSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
