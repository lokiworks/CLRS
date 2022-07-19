package main

func plusOne(digits []int) []int {
	l := len(digits)
	remainder := 0
	var result []int

	for i := l - 1; i >= 0; i-- {
		var temp int
		if i == l-1 {
			temp = (remainder + digits[i] + 1) % 10
			remainder = (remainder + digits[i] + 1) / 10
		} else {
			temp = (remainder + digits[i]) % 10
			remainder = (remainder + digits[i]) / 10
		}
		result = append(result, temp)
	}
	if remainder != 0 {
		result = append(result, remainder)
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
