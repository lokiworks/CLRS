package main

import "testing"

func Test_minimum(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "WithMin1",
			args: args{a: []int{5, 3, 1, 2}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimum(tt.args.a); got != tt.want {
				t.Errorf("minimum() = %v, want %v", got, tt.want)
			}
		})
	}
}
