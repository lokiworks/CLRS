package main

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"WithLongestCommonPrefixAsFL",
			args{strs: []string{"flower", "flow", "flight"}},
			"fl",
		},
		{
			"WithNoLongestCommonPrefix",
			args{strs: []string{"dog", "rarecar", "car"}},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
