package main

import "testing"

func Test_randomizedSelect(t *testing.T) {
	type args struct {
		a []int
		p int
		r int
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "With4NthAs6",
			args: args{a: []int{6, 19, 4, 12, 14, 9, 15, 7, 8, 11, 3, 13, 2, 5, 10}, p: 0, r: len([]int{6, 19, 4, 12, 14, 9, 15, 7, 8, 11, 3, 13, 2, 5, 10}) - 1, i: 4},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randomizedSelect(tt.args.a, tt.args.p, tt.args.r, tt.args.i); got != tt.want {
				t.Errorf("randomizedSelect() = %v, want %v", got, tt.want)
			}
		})
	}
}
