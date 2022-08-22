package main

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

func myAtoi(s string) int {

	var sb strings.Builder
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	if (s[0] != '-' && s[0] != '+') && !unicode.IsDigit(rune(s[0])) {
		return 0
	}
	for i, e := range s {
		if i == 0 {
			sb.WriteByte(byte(e))
			continue
		}

		if unicode.IsDigit(e) {
			sb.WriteByte(byte(e))

		}
		if e == '.' || unicode.IsLetter(e) || unicode.IsSpace(e) || (e == '+' || e == '-') {
			break
		}
	}
	r, _ := strconv.Atoi(strings.TrimSpace(sb.String()))
	if r < 0 {
		if r < math.MinInt32 {
			return math.MinInt32
		}
	} else {
		if r > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return r
}

func myAtoiV2(s string) int {
	// 取出空格
	s = strings.TrimSpace(s)

	jump := func(e rune) bool {
		if e == '.' || unicode.IsLetter(e) || unicode.IsSpace(e) || e == '+' || e == '-' {
			return true
		}
		return false
	}
	r := 0
	signed := 1

	for i, e := range s {
		if i == 0 {
			if e == '-' {
				signed = -1
			} else if e == '+' {
				signed = 1
			} else if unicode.IsDigit(e) {
				r = 10*r + int(e-'0')
			} else {
				break
			}
			continue
		}

		if jump(e) {
			break
		}
		if (r * signed) > math.MaxInt32/10 {
			return math.MaxInt32
		}

		if (r * signed) < math.MinInt32/10 {
			return math.MinInt32
		}

		r = 10*r + int(e-'0')

	}

	if signed > 0 {
		if r > math.MaxInt32 {
			return math.MaxInt32
		}
	} else {
		if signed*r < math.MinInt32 {
			return math.MinInt32
		}
	}

	return signed * r
}
