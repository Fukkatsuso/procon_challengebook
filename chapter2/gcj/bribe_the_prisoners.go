package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initialBufSize = 100000
	maxBufSize     = 10000000
	inf            = 1 << 60
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wt = bufio.NewWriter(os.Stdout)
)

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func getInt2() (int, int) {
	return getInt(), getInt()
}

func getInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = getInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	p, q := getInt2()
	a := make([]int, q+2)
	a[0], a[q+1] = 0, p+1 // 番兵
	for i := 1; i <= q; i++ {
		a[i] = getInt()
	}

	dp := make([][]int, p+2)
	for i := range dp {
		dp[i] = make([]int, q+2)
		for j := range dp[i] {
			if abs(i-j) > 1 {
				dp[i][j] = inf
			}
		}
	}

	for w := 2; w <= q+1; w++ {
		for i := 0; i+w <= q+1; i++ {
			j := i + w
			minCost := inf
			for k := i + 1; k < j; k++ {
				minCost = min(minCost, dp[i][k]+dp[k][j])
			}
			dp[i][j] = minCost + a[j] - a[i] - 2
		}
	}

	puts(dp[0][q+1])
}
