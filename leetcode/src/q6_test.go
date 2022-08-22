package main

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"WithOneRow",
			args{s: "A", numRows: 1},
			"A",
		},

		{
			"WithThreeRows",
			args{s: "PAYPALISHIRING", numRows: 3},
			"PAHNAPLSIIGYIR",
		},
		{
			"WithFourRows",
			args{s: "PAYPALISHIRING", numRows: 4},
			"PINALSIGYAHRPI",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
