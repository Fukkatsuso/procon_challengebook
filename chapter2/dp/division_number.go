package main

import "fmt"

func main() {
	var n, m, M int
	fmt.Scan(&n, &m, &M)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	dp[0][0] = 1
	for i := 1; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if j-i >= 0 {
				dp[i][j] = (dp[i][j-i] + dp[i-1][j]) % M
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	fmt.Println(dp[m][n])
}

// 4 3 10000

//-> 4
