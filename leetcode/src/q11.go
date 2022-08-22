package main

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	maxArea := 0

	for l < r {
		area := min(height[l], height[r]) * (r - l)
		maxArea = max(maxArea, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxArea

}
