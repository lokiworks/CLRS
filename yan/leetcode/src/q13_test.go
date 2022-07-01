package main

import "testing"

func Test_romanToInt(t *testing.T) {
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
			"WithRomanIII",
			args{s: "III"},
			3,
		},
		{
			"WithRomanLVIII",
			args{s: "LVIII"},
			58,
		},

		{
			"WithRomanMCMXCIV",
			args{s: "MCMXCIV"},
			1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
