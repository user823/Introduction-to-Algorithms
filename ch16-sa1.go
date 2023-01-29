package main

import (
	"fmt"
)

func main() {
	//活动已经按结束时间非递减排序
	//贪心选择：每次选择最早结束的活动
	var s = []int{1, 3, 0, 5, 3, 5, 6, 8, 8, 2, 12}
	var f = []int{4, 5, 6, 7, 9, 9, 10, 11, 12, 14, 16}
	fmt.Println(recursiceActivitySelect(s, f, -1, 11))
}

func recursiceActivitySelect(s []int, f []int, k int, n int) []int{  //-1表示初始虚拟活动
	m := k + 1
	if m == 0 {
		var ans = []int{0}
		return append(ans, recursiceActivitySelect(s, f, m, n)...)
	} 
	for m < n && s[m] < f[k] {
		m++
	}
	if m < n {
		var ans = []int{m}
		return append(ans, recursiceActivitySelect(s, f, m, n)...)
	}
	return nil
}
