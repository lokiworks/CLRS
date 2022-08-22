package main

import "fmt"

const OpenAddressArraySize = 16

func openAddressHashInsert(T []int, k int) int {
	i := 0
	for i < len(T) {
		q := openAddressHash(k, i)
		if T[q] == 0 {
			T[q] = k
			return q
		} else {
			i++
		}

	}
	return 0

}

func openAddressHashSearch(T []int, k int) int {
	i := 0
	for i < len(T) {
		q := openAddressHash(k, i)
		if q == 0 {
			return 0
		}
		if T[q] == k {
			return q
		} else {
			i++
		}
	}
	return 0
}

func openAddressHash(key int, i int) int {
	return ((key % OpenAddressArraySize) + i) % OpenAddressArraySize
}

func main() {
	array := make([]int, OpenAddressArraySize)
	openAddressHashInsert(array, 1)
	openAddressHashInsert(array, 2)
	openAddressHashInsert(array, 3)

	fmt.Println(openAddressHashSearch(array, 1))
	fmt.Println(openAddressHashSearch(array, 2))
	fmt.Println(openAddressHashSearch(array, 3))
	fmt.Println(openAddressHashSearch(array, 4))
}
