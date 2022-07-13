package main

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"WithPositiveNumber",
			args{s: "42"},
			42,
		},
		{
			"WithSingedPositiveNumber",
			args{s: "+42"},
			42,
		},
		{
			"WithSpaceAndNegativeNumber",
			args{s: "   -42"},
			-42,
		},
		{
			"WithNumberAndSpaceAndLetter",
			args{s: "4193 with words"},
			4193,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myAtoiV2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"WithPositiveNumber",
			args{s: "42"},
			42,
		},
		{
			"WithSingedPositiveNumber",
			args{s: "+42"},
			42,
		},
		{
			"WithSpaceAndNegativeNumber",
			args{s: "   -42"},
			-42,
		},
		{
			"WithNumberAndSpaceAndLetter",
			args{s: "4193 with words"},
			4193,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoiV2(tt.args.s); got != tt.want {
				t.Errorf("myAtoiV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
