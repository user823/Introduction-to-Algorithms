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
	if len1 < len2 {
		return help(x, len1, y, len2)
	}
	return help(y, len2, x, len1)
}

func help(x []int, len1 int, y []int, len2 int) int {      //x为短，y为长
	dp := make([]int, len1)
	for i:=0; i<len2; i++ {   
		for j:=0; j<len1; j++{
			if i==0 || j==0 {
				dp[j] = 0
			} else if y[i] == x[j] {
				dp[j] = dp[j-1] + 1
			} else {
				dp[j] = max(dp[j-1], dp[j])
			}
		}
	}
	return dp[len1-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
