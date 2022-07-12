package main

import "math"

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}
	r := 0

	for x != 0 {
		r = 10*r + x%10
		x = x / 10
	}
	if sign < 0 {
		if sign*r < int(math.Pow(-2, 31)) {
			return 0
		}
	} else {
		if r > int(math.Pow(2, 31))-1 {
			return 0
		}
	}
	return sign * r
}

func reverseV2(x int) int {
	r := 0
	for x != 0 {
		if r > math.MaxInt32/10 {
			return 0
		}

		if r < math.MinInt32/10 {
			return 0
		}
		r = 10*r + x%10
		x = x / 10
	}
	return r
}

func main() {
	println(reverseV2(-2147483412))
	println(reverse(123))
	println(reverse(120))
	println(reverse(-123))

}
