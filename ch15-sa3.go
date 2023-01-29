package main

import (
	"fmt"
)

func main() {
	fmt.Println(getLCS([]int{1, 3, 4, 5, 7}, []int{2, 3, 5, 8}))
}

func getLCS(x []int, y []int) int {
	len1 := len(x)
	len2 := len(y)
	dp := make([][]int, len1)
	for i:=0; i<len1; i++ {
		dp[i] = make([]int, len2)
		for j:=0; j<len2; j++ {
			if i==0 || j==0 {
				dp[i][j] = 0
			} else if x[i] == y[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len1-1][len2-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
