package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var M int
	fmt.Scan(&M)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		dp[i][0] = 1
	}

	for i := 0; i < n; i++ {
		for j := 1; j <= m; j++ {
			if j-1 >= a[i] {
				dp[i+1][j] = (dp[i][j] + dp[i+1][j-1] - dp[i][j-1-a[i]]) % M
			} else {
				dp[i+1][j] = (dp[i][j] + dp[i+1][j-1]) % M
			}
		}
	}

	fmt.Println(dp[n][m])
}

// 3 3
// 1 2 3
// 10000

//-> 6
