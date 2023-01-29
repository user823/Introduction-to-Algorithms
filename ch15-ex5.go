package main

import (
	"fmt"
)

func main() {
	fmt.Println(getLS([]int{1, 3, 4, 0, 7}))
}


func getLS(nums []int) int {
	length := len(nums)
	var dp = make([]int, length)
	dp[0] = 1
	for i:=1; i<length; i++ {
		dp[i] = 1
		for j:=0; j<i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return dp[length-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
