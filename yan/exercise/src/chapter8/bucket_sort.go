package main

import "math"

func bucketSort(a []float64) []float64 {
	n := len(a)
	b := make([][]float64, n)
	for i := 0; i < n; i++ {
		b[i] = []float64{}
	}

	for _, e := range a {
		b[int(math.Floor(float64(n)*e))] = append(b[int(math.Floor(float64(n)*e))], e)
	}
	for i := 0; i < n; i++ {
		insertSort(b[i])
	}

	c := []float64{}
	for i := 0; i < n; i++ {
		c = append(c, b[i]...)
	}
	copy(a, c)
	return a

}

func insertSort(a []float64) {
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

}

func main() {
	a := bucketSort([]float64{0.897, 0.565, 0.656, 0.1234})
	println(a)
}
