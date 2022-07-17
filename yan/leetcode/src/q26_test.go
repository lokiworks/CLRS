package main

import (
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		// TODO: Add test cases.
		{
			name:  "WithOnePlaceRepeat",
			args:  args{[]int{1, 1, 2}},
			want:  2,
			want1: []int{1, 2},
		},
		{
			name:  "WithFourPlaceRepeat",
			args:  args{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			want:  5,
			want1: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := removeDuplicates(tt.args.nums)
			if got != tt.want {
				t.Errorf("removeDuplicates() got = %v, want %v", got, tt.want)
			}
			for i := 0; i < tt.want; i++ {
				if got1[i] != tt.want1[i] {
					t.Errorf("removeDuplicates() got1[%v] = %v, want %v", i, got1[i], tt.want1[i])

				}
			}
		})
	}
}
