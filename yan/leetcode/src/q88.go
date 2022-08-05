package main

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	a := make([]int, m+n)
	i, j, k := 0, 0, 0

	for i < m && j < n {
		if nums1[i] < nums2[j] {
			a[k] = nums1[i]
			i++
		} else {
			a[k] = nums2[j]
			j++
		}
		k++
	}
	for ; i < m; i, k = i+1, k+1 {
		a[k] = nums1[i]
	}
	for ; j < n; j, k = j+1, k+1 {
		a[k] = nums2[j]
	}

	copy(nums1, a)

	return a
}
