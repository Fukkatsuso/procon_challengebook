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
	INF            = 1 << 60
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

func getInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = getInt()
	}
	return slice
}

func getFloat64() float64 {
	f, _ := strconv.ParseFloat(gets(), 64)
	return f
}

func getFloat64s(n int) []float64 {
	slice := make([]float64, n)
	for i := 0; i < n; i++ {
		slice[i] = getFloat64()
	}
	return slice
}

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func min(nums ...float64) float64 {
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

	n, m := getInt(), getInt()
	a, b := getInt()-1, getInt()-1
	t := getInts(n)
	// グラフの距離行列
	d := make([][]int, m)
	for i := range d {
		d[i] = make([]int, m)
		for j := range d[i] {
			d[i][j] = INF
		}
	}
	e := getInt()
	for i := 0; i < e; i++ {
		u, v, cost := getInt()-1, getInt()-1, getInt()
		d[u][v], d[v][u] = cost, cost
	}

	dp := make([][]float64, 1<<n)
	for i := range dp {
		dp[i] = make([]float64, m)
		for j := range dp[i] {
			dp[i][j] = float64(INF)
		}
	}
	dp[(1<<n)-1][a] = 0

	res := float64(INF)
	for S := (1 << n) - 1; S >= 0; S-- {
		res = min(res, dp[S][b])
		for i := 0; i < n; i++ {
			// 乗車券iが使用済み
			if (S>>i)&1 == 0 {
				continue
			}
			// 乗車券iを使って都市uから都市vへ移動
			for u := 0; u < m; u++ {
				for v := 0; v < m; v++ {
					dp[S&^(1<<i)][v] = min(
						dp[S&^(1<<i)][v],
						dp[S][u]+float64(d[u][v])/float64(t[i]),
					)
				}
			}
		}
	}

	if res == INF {
		puts("Impossible")
	} else {
		putf("%.3f\n", res)
	}
}
