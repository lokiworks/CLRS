package main

func select1(a []int, p, r, i int) int {
	for (r-p)%5 != 0 {
		for j := p + 1; j < r+1; j++ {
			if a[p] > a[j] {
				a[p], a[j] = a[j], a[p]
			}
		}
		if i == 0 {
			return a[p]
		}
		p++
		i--
	}
	g := (r - p) / 5
	for j := p; j < p+g; j++ {

		a2 := []int{a[j], a[j+g], a[j+2*g], a[j+3*g], a[j+4*g]}
		insertSort(a2)
		a[j], a[j+g], a[j+2*g], a[j+3*g], a[j+4*g] = a2[0], a2[1], a2[2], a2[3], a2[4]

	}
	x := select1(a, p+2*g, p+3*g-1, g/2)
	q := partition2(a, p, r, x)
	k := q - p + 1
	if i == k {
		return a[q]
	} else if i < k {
		return select1(a, p, q-1, i)
	}
	return select1(a, q+1, r, i-k)
}

func partition2(a []int, p, r, x int) int {
	for i := p; i < r; i++ {
		if a[i] == x {
			a[i], a[r] = a[r], a[i]
			break
		}
	}

	return partition(a, p, r)
}

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

func main() {
	v := select1([]int{6, 19, 4, 12, 14, 9, 15, 7, 8, 11, 3, 13, 2, 5, 10}, 0, 14, 4)
	println(v)
}
