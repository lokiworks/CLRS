package main

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"WithFullMatch",
			args{s: "aa", p: "a"},
			false,
		},
		{
			"WithZeroTimesMatch",
			args{s: "c", p: "ca*"},
			true,
		},
		{
			"WithOneTimesMatch",
			args{s: "ca", p: "ca*"},
			true,
		},
		{
			"WithTwoTimesMatch",
			args{s: "aa", p: "a*"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isMatchV2(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"WithFullMatch",
			args{s: "aa", p: "a"},
			false,
		},
		{
			"WithZeroTimesMatch",
			args{s: "c", p: "ca*"},
			true,
		},
		{
			"WithOneTimesMatch",
			args{s: "ca", p: "ca*"},
			true,
		},
		{
			"WithTwoTimesMatch",
			args{s: "aa", p: "a*"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatchV2(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatchV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
