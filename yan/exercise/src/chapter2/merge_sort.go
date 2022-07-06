package chaptor2

func innerMergeSort(a []int, p int, r int) []int {
	if p >= r {
		return a
	}
	q := (p + r) / 2
	innerMergeSort(a, p, q)
	innerMergeSort(a, q+1, r)
	// Merge a[p:q] and a[q+1:r] into a[p:r]
	merge(a, p, q, r)
	return a
}

func mergeSort(a []int) []int {
	return innerMergeSort(a, 0, len(a)-1)
}

func merge(a []int, p int, q int, r int) {
	lLen := q - p + 1 // length of a[p:q]
	rLen := r - q     // lenght of a[q+1:r]
	lArray := make([]int, lLen)
	rArray := make([]int, rLen)

	// copy a[p:q] into lArray[0:lLen-1]
	for i := 0; i < lLen; i++ {
		lArray[i] = a[p+i]
	}
	// copy a[q+1:r] into rArray[0:rLen-1]
	for i := 0; i < rLen; i++ {
		rArray[i] = a[q+i+1]
	}

	i := 0 // i indexes the smallest remaining element in lArray
	j := 0 // j indexes the smallest remaining element in rArray
	k := p // k indexes the location in a to fill

	// as long as each of the arrays lArray and rArray contains an unmerged element,
	// copy the smallest unmerged element back into a[p:r]
	for i < lLen && j < rLen {
		if lArray[i] <= rArray[j] {
			a[k] = lArray[i]
			i += 1
		} else {
			a[k] = rArray[j]
			j += 1
		}
		k++
	}
	// having gone through one of lArray and rArray entirely, copy the
	// remainder of the other to the end of a[p:r]
	for i < lLen {
		a[k] = lArray[i]
		i++
		k++
	}
	for j < rLen {
		a[k] = rArray[j]
		j++
		k++
	}
}
