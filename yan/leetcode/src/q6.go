package main

import (
	"strings"
)

// Need Again!!!
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	l := len(s)
	cycleLen := 2 * (numRows - 1)
	var sb = strings.Builder{}

	for i := 0; i < numRows; i++ {
		for j := 0; i+j < l; j += cycleLen {
			sb.WriteByte(s[i+j])
			if i != 0 && i != numRows-1 && (cycleLen+j-i) < l {
				sb.WriteByte(s[cycleLen+j-i])
			}
		}
	}
	return sb.String()

}
