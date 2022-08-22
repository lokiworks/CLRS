package main

func minimum(a []int) int {
	n := len(a)
	if len(a) == 0 {
		return 0
	}

	min := a[0]

	for i := 1; i < n; i++ {
		if min > a[i] {
			min = a[i]
		}
	}

	return min
}
