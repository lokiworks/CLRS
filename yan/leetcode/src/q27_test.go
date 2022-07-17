package main

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "WithOnePlaceRepeat",
			args: args{[]int{1, 1, 2}, 2},
			want: 2,
		},
		{
			name: "WithFourPlaceRepeat",
			args: args{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 1},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
