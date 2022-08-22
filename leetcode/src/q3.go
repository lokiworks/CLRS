package main

import (
	"math"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	s = strings.TrimSpace(s)
	l = len(s)
	if l == 0 || l == 1 {
		return 1
	}

	m := 0
	subMap := make(map[byte]int)
	for i := 0; i < l; i++ {
		subMap = map[byte]int{}
		subMap[s[i]] = i
		for j := i + 1; j < l; j++ {
			_, ok := subMap[s[j]]
			if ok {
				break
			} else {
				subMap[s[j]] = j
			}
		}
		if m < len(subMap) {
			m = len(subMap)
		}
	}
	return m
}

func lengthOfLongestSubstringV2(s string) int {
	if len(s) == 0 {
		return 0
	}

	m := make(map[byte]int)
	left := 0
	max := 0
	for i, e := range s {
		v, ok := m[byte(e)]
		if ok {
			left = int(math.Max(float64(left), float64(v+1)))
		}
		m[byte(e)] = i
		max = int(math.Max(float64(max), float64(i-left+1)))

	}
	return max
}
