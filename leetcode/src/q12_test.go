package main

import "testing"

func Test_intToRomanV2(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "With3",
			args: args{num: 3},
			want: "III",
		},
		{
			name: "With58",
			args: args{num: 58},
			want: "LVIII",
		},
		{
			name: "With1994",
			args: args{num: 1994},
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRomanV2(tt.args.num); got != tt.want {
				t.Errorf("intToRomanV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "With3",
			args: args{num: 3},
			want: "III",
		},
		{
			name: "With58",
			args: args{num: 58},
			want: "LVIII",
		},
		{
			name: "With1994",
			args: args{num: 1994},
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
