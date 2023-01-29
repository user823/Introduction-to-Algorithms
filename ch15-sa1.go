package main

import (
	"fmt"
)

func main() {
	for i:=1; i<=10; i++ {
		fmt.Println(getBiggest(i))
	}
}

func getBiggest(n int) int {
	var tb = []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	var dp = make([][]int, n+1)      //定义二维切片
	dp[0] = make([]int, 11)
	for i:=1; i<=n; i++ {            //i表示当前长度
		dp[i] = make([]int, 11)
		max := 0
		for j:=1; j<=i; j++ {
			dp[i][j] = tb[j] + dp[i-j][0]
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
		dp[i][0] = max
	}
	return dp[n][0]
}
