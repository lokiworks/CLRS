package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, e := range nums {
		m[e] = i
	}

	for i, e := range nums {
		r := target - e
		j, ok := m[r]
		if ok {
			if i != j {
				return []int{i, j}
			}

		}
	}
	return nil
}
