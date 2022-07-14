package main

func parent(i int) int {
	return (i - 1) / 2
}
func left(i int) int {
	return 2*(i+1) - 1
}

func right(i int) int {
	return 2 * (i + 1)
}

func maxHeapify(a []int, heapSize, i int) {
	l := left(i)
	r := right(i)
	largestIndex := -1
	if l < heapSize {
		if a[l] > a[i] {
			largestIndex = l
		} else {
			largestIndex = i
		}
	}
	if r < heapSize {
		if a[r] > a[largestIndex] {
			largestIndex = r
		}
	}
	if largestIndex != -1 && largestIndex != i {
		// exchange a[i] with a[]
		a[i], a[largestIndex] = a[largestIndex], a[i]
		maxHeapify(a, heapSize, largestIndex)
	}

}

func buildMaxHeap(a []int) {
	l := len(a) - 1
	for i := l / 2; i >= 0; i-- {
		maxHeapify(a, l, i)
	}
}

func heapSort(a []int) []int {
	buildMaxHeap(a)
	l := len(a)
	for i := len(a) - 1; i >= 1; i-- {
		a[0], a[i] = a[i], a[0]
		l--
		maxHeapify(a, l, 0)

	}
	return a
}
