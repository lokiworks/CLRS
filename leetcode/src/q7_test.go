package main

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"WithPositiveNumber",
			args{x: 123},
			321,
		},
		{
			"WithNegativeNumber",
			args{x: -123},
			-321,
		},
		{
			"WithEndWithZero",
			args{x: 120},
			21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseV2(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"WithPositiveNumber",
			args{x: 123},
			321,
		},
		{
			"WithNegativeNumber",
			args{x: -123},
			-321,
		},
		{
			"WithEndWithZero",
			args{x: 120},
			21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseV2(tt.args.x); got != tt.want {
				t.Errorf("reverseV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
