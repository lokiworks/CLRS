package main

func quickSort(a []int) []int {
	innerQuickSort(a, 0, len(a)-1)
	return a
}

func innerQuickSort(a []int, p, r int) {
	if p < r {
		// partition the subarray around the pivot, which ends up in A[q].
		q := partition(a, p, r)
		innerQuickSort(a, p, q-1) // recursively sort the low side
		innerQuickSort(a, q+1, r) // recursively sort the high side
	}
}

func partition(a []int, p, r int) int {
	x := a[r]                // the pivot
	i := p - 1               // the highest index into the low side
	for j := p; j < r; j++ { // process each element other than the pivot
		if a[j] <= x { // does this element belong on the low side?
			i++ //index of a new slot in the low side
			a[i], a[j] = a[j], a[i]

		}

	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1

}
