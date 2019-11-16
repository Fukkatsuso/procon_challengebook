package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	w := make([]int, n)
	v := make([]int, n)
	vMax := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&w[i])
		fmt.Scan(&v[i])
		if vMax < v[i] {
			vMax = v[i]
		}
	}
	var W int
	fmt.Scan(&W)
	VMax := n * vMax
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, VMax+1)
		for j := 0; j <= VMax; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0

	for i := 0; i < n; i++ {
		var a, b int
		for j := 0; j <= VMax; j++ {
			a = dp[i][j]
			if j < v[i] {
				dp[i+1][j] = a
				continue
			}
			b = dp[i][j-v[i]] + w[i]
			if a <= b {
				dp[i+1][j] = a
			} else {
				dp[i+1][j] = b
			}
		}
	}

	ans := 0
	for i := 0; i <= VMax; i++ {
		if W >= dp[n][i] {
			ans = i
		}
	}
	fmt.Println(ans)
}

// 4
// 2 3
// 1 2
// 3 4
// 2 2
// 5

//-> 7
