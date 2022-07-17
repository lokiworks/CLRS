package main

func searchInsert(nums []int, target int) int {
	mv := -1
	for i, e := range nums {
		if e == target {
			return i
		} else if e < target {
			mv = i
		}
	}
	return mv + 1
}
