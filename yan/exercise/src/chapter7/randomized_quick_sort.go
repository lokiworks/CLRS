package main

import (
	"math/rand"
	"time"
)

func init() {
	// set seed
	rand.Seed(time.Now().UnixNano())
}

func randomizedPartition(a []int, p, r int) int {
	i := rand.Intn(r-p) + p
	a[i], a[r] = a[r], a[i]
	return partition(a, p, r)

}

func innerRandomizedQuickSort(a []int, p, r int) {
	if p < r {
		q := randomizedPartition(a, p, r)
		innerRandomizedQuickSort(a, p, q-1)
		innerRandomizedQuickSort(a, q+1, r)
	}
}

func randomizedQuickSort(a []int) []int {
	innerRandomizedQuickSort(a, 0, len(a)-1)
	return a
}

func main() {
	a := randomizedQuickSort([]int{5, 4, 2, 2, 1})
	println(a)
}
