package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
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
			"withLength4",
			args{s: "abcaddcbb"},
			4,
		},
		{
			"withOneChar",
			args{s: "c"},
			1,
		},
		{
			"withEmptyString",
			args{s: ""},
			0,
		},
		{
			"withSpaceString",
			args{s: " "},
			1,
		},
		{
			"withTwoChar",
			args{s: "cu"},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLongestSubstringV2(t *testing.T) {
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
			"withLength4",
			args{s: "abcaddcbb"},
			4,
		},
		{
			"withOneChar",
			args{s: "c"},
			1,
		},
		{
			"withEmptyString",
			args{s: ""},
			0,
		},
		{
			"withSpaceString",
			args{s: " "},
			1,
		},
		{
			"withTwoChar",
			args{s: "cu"},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringV2(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
