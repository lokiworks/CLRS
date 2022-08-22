package main

import "strings"

func lengthOfLastWord(s string) int {

	sub := strings.Split(strings.TrimSpace(s), " ")
	if len(sub) == 0 {
		return 0
	}
	return len(sub[len(sub)-1])
}
