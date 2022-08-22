package main

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "WithN2",
			args: args{n: 2},
			want: 2,
		},
		{
			name: "WithN3",
			args: args{n: 3},
			want: 3,
		},
		{
			name: "WithN4",
			args: args{n: 4},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbStairsV2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "WithN2",
			args: args{n: 2},
			want: 2,
		},
		{
			name: "WithN3",
			args: args{n: 3},
			want: 3,
		},
		{
			name: "WithN4",
			args: args{n: 4},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairsV2(tt.args.n); got != tt.want {
				t.Errorf("climbStairsV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
