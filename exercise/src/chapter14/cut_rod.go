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

func main() {
	pArray := []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	p := cut_rod(pArray, 3)
	fmt.Println(p)
}
