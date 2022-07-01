package main

import "strings"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 1
	}
	l := len(s)

	if l == 1 {
		return 1
	}

	m := 0
	for i := 0; i < l; i++ {
		subMap := make(map[rune]int)
		subMap[rune(s[i])] = i
		for j := i + 1; j < l; j++ {
			_, ok := subMap[rune(s[j])]
			if ok {
				break
			} else {
				subMap[rune(s[j])] = j
			}
		}
		if m < len(subMap) {
			m = len(subMap)
		}
	}
	return m
}
