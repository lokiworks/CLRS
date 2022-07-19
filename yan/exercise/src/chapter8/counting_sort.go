package main

func countingSort(a []int, k int) []int {
	return innerCountingSort(a, len(a), k)
}

func innerCountingSort(a []int, n, k int) []int {
	r := make([]int, len(a))
	c := make([]int, k+1)

	for _, e := range a {
		c[e]++
	}

	for i := 1; i < k+1; i++ {
		c[i] += c[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		r[c[a[i]]-1] = a[i]
		c[a[i]]--
	}

	return r
}
