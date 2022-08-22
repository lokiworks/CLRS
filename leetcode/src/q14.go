package main

import "strings"

func longestCommonPrefix(strs []string) string {

	longestStr := ""

	for i := 0; i < len(strs[0]); i++ {
		longestStr = strs[0][:i+1]
		for _, str := range strs {
			if !strings.HasPrefix(str, longestStr) {
				return strs[0][:i]
			}
		}
	}

	return longestStr
}
