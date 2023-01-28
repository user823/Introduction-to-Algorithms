package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getBST(5, []float64{0, 0.15, 0.10, 0.05, 0.10, 0.20}, []float64{0.05, 0.10, 0.05, 0.05, 0.05, 0.10}))
}

func getBST(n int, p []float64, q []float64) float64{
	var e = make([][]float64, n+2) //e[1...n+1, 0...n]
	var w = make([][]float64, n+2) //w[1...n+1, 0...n]

	//初始化
	for i:=1; i<n+2; i++{
		e[i] = make([]float64, n+1)
		w[i] = make([]float64, n+1)
		e[i][i-1] = q[i-1]
		w[i][i-1] = q[i-1]
	}

	//以长度为树的规模对表进行计算 
	//长度取值为1...n 
	for l:=1; l<=n; l++ { 
		for i:=1; i<=n+1-l; i++ {
			j := i - 1 + l
			e[i][j] = math.MaxFloat64
			w[i][j] = w[i][j-1] + p[j] + q[j]
			for r:=i; r<=j; r++ {
				t := e[i][r-1] + e[r+1][j] + w[i][j]
				if t < e[i][j] {
					e[i][j] = t
				}
			}
		}
	}
	return e[1][n]
}
