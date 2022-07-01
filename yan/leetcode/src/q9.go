package main

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.FormatInt(int64(x), 10)

	l := len(s)
	for i := 0; i < l/2+1; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true

}
