package main

import "math"

func radixSort(a []int, d int) []int {
	return innerRadixSort(a, len(a), d)
}

func innerRadixSort(a []int, n, d int) []int {

	for i := 0; i < d; i++ {
		countingSortForRadix(a, n, int(math.Pow10(i)))

	}
	return a
}

func countingSortForRadix(a []int, n, k int) {
	r := make([]int, len(a))
	c := make([]int, 10)

	for _, e := range a {
		c[(e/k)%10]++
	}

	for i := 1; i < 10; i++ {
		c[i] += c[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		r[c[(a[i]/k)%10]-1] = a[i]
		c[(a[i]/k)%10]--
	}

	copy(a, r)
}
