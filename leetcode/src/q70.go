package main

// 3
// 1*i+2*j
// i = 0
// i = 1 j = 1

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// v2: 滚动数组版本
func climbStairsV2(n int) int {
	first := 1
	second := 2
	for i := 3; i < n+1; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}
