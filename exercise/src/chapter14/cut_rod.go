package main

import (
	"fmt"
	"math"
)

func cut_rod(p []int, n int) int {
	if n == 0 {
		return 0
	}
	q := math.MinInt
	for i := 0; i < n; i++ {
		q = int(math.Max(float64(q), float64(p[i]+cut_rod(p, n-i-1))))
	}
	return q
}

func cut_rod_v2(p []int, n int) int {
	if n == 0 {
		return 0
	}
	r := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		r[i] = math.MinInt
	}

	return memorized_cut_rod_aux(p, n, r)
}

func memorized_cut_rod_aux(p []int, n int, r []int) int {
	if r[n] >= 0 {
		return r[n]
	}
	var q int
	if n == 0 {
		q = 0
	} else {
		q = math.MinInt
		for i := 0; i < n; i++ {
			q = int(math.Max(float64(q), float64(p[i]+memorized_cut_rod_aux(p, n-i-1, r))))
		}
	}
	r[n] = q

	return q
}

// bottom-up-cut-rod
func cut_rod_v3(p []int, n int) int {
	r := make([]int, n+1)
	for j := 1; j < n+1; j++ {
		q := math.MinInt
		for i := 1; i <= j; i++ {
			q = int(math.Max(float64(q), float64(p[i-1]+r[j-i])))
		}
		r[j] = q
	}
	return r[n]
}

func main() {
	pArray := []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	p := cut_rod(pArray, 3)
	fmt.Println(p)

	p = cut_rod_v2(pArray, 3)
	fmt.Println(p)

	p = cut_rod_v3(pArray, 3)
	fmt.Println(p)

}
