package main

import (
	"regexp"
)

func isMatch(s string, p string) bool {
	m, _ := regexp.MatchString("^"+p+"$", s)
	return m
}

// Need Again!!!
func isMatchV2(s string, p string) bool {
	m, n := len(s), len(p)
	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}
	// 以下标0打头
	// 表示s为空串
	// 表示p为空串
	f[0][0] = true
	for i := 0; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}
