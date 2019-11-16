// x
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	m := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&m[i])
	}
	var K int
	fmt.Scan(&K)
	dp := make([]int, K+1)
	for i := 0; i < K+1; i++ {
		dp[i] = -1
	}
	dp[0] = 0

	for i := 0; i < n; i++ {
		for j := 0; j <= K; j++ {
			if dp[j] >= 0 {
				dp[j] = m[i]
			} else if j < a[i] || dp[j-a[i]] <= 0 {
				dp[j] = -1
			} else {
				dp[j] = dp[j-a[i]] - 1
			}
		}
	}

	if dp[K] >= 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

// 3
// 3 5 8
// 3 2 2
// 17

//-> Yes
