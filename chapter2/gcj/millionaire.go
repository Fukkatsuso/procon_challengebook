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

func getFloat64() float64 {
	f, _ := strconv.ParseFloat(gets(), 64)
	return f
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func max(nums ...float64) float64 {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	m := getInt()
	p := getFloat64()
	x := getInt()

	n := 1 << m

	// dp[r][i]: 残りrラウンド時点で所持金(i*1000000/2^i)以上((i+1)*1000000/2^i)未満だった場合の成功確率
	dp := make([][]float64, m+1)
	for i := range dp {
		dp[i] = make([]float64, n+1)
	}
	dp[0][n] = 1.0

	for r := 0; r < m; r++ {
		for i := 0; i <= n; i++ {
			t := 0.0
			for j := 0; i+j <= n && i-j >= 0; j++ {
				t = max(t, p*dp[r][i+j]+(1-p)*dp[r][i-j])
			}
			dp[r+1][i] = t
		}
	}

	i := x * n / 1000000
	ans := dp[m][i]
	puts(ans)
}
