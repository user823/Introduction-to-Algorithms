package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getLowest(6))
}

func getLowest(n int) int {        //n表示矩阵数目
	var tb = []int{30, 35, 15, 5, 10, 20, 25} //A1 ... A6
	var dp = make([][]int, 6)      //dp表示矩阵相乘的代价

	//初始化
	for i:=0; i<6; i++ {
		dp[i] = make([]int, 6)
		dp[i][i] = 0
		if i < 5 {
			dp[i][i+1] = tb[i] * tb[i+1] * tb[i+2]
		}
	}

	for i:=n-2; i>=0 ;i-- {      
		for j:=i+2; j<n; j++ {
			min := math.MaxInt
			for k:=i; k<j; k++ {
				tmp := dp[i][k] + dp[k+1][j] + tb[i] * tb[k+1] * tb[j+1]
				if tmp < min {
					min = tmp
				}
			}
			dp[i][j] = min
		}
	}
	
	return dp[0][n-1]
}
