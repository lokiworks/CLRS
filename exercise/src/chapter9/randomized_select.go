package main

import "math/rand"

func randomizedSelect(a []int, p, r, i int) int {
	if p == r {
		return a[p]
	}
	q := randomizedPartition(a, p, r)
	k := q - p
	if i == k {
		return a[q]
	} else if i < k {
		return randomizedSelect(a, p, q-1, i)
	} else {
		return randomizedSelect(a, q+1, r, i-k-1)
	}
}

func randomizedPartition(a []int, p, r int) int {
	i := rand.Intn(r-p) + p
	a[i], a[r] = a[r], a[i]
	return partition(a, p, r)

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
