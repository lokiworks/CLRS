package main

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"WithOneRightBracket",
			args{s: "]"},
			false,
		},
		{
			"WithTwoMatchBracket",
			args{s: "{[]}"},
			true,
		},
		{
			"WithTwoNotMatchBracket",
			args{s: "([)]"},
			false,
		},
		{
			"WithThreeMatchBracket",
			args{s: "()[]{}"},
			true,
		},
		{
			"WithOneNotMatchBracket",
			args{s: "(]"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
