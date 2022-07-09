package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lLen := len(nums1)
	rLen := len(nums2)
	prev := 0
	cur := 0
	len := (lLen+rLen)/2 + 1
	lIndex := 0
	rIndex := 0
	for i := 0; i < len; i++ {
		prev = cur
		if lIndex < lLen && (rIndex >= rLen || nums1[lIndex] < nums2[rIndex]) {
			cur = nums1[lIndex]
			lIndex++
		} else {
			cur = nums2[rIndex]
			rIndex++
		}

	}
	if (lLen+rLen)%2 == 0 {
		return float64(cur+prev) / 2
	} else {
		return float64(cur)
	}
}

func main() {
	println(findMedianSortedArrays([]int{0, 0}, []int{0, 0}))
}
