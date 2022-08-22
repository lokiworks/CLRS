package main

func removeElement(nums []int, val int) int {
	var a []int
	for _, e := range nums {
		if e != val {
			a = append(a, e)
		}
	}

	for i := 0; i < len(a); i++ {
		nums[i] = a[i]
	}

	return len(a)

}
