package main

func removeDuplicates(nums []int) (int, []int) {
	m := make(map[int]int)
	var a []int
	for i, e := range nums {
		_, ok := m[e]
		if !ok {
			m[e] = i
			a = append(a, e)
		}
	}

	for i := 0; i < len(m); i++ {
		nums[i] = a[i]
	}

	return len(m), nums
}
