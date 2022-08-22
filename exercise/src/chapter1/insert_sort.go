package main

func insertSort(a []int) []int {
	len := len(a)
	for i := 1; i < len; i++ {
		k := a[i]
		// insert k into the sorted subarray A[0:i-1]
		j := i - 1
		for j >= 0 && a[j] > k {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = k
	}

	return a
}
