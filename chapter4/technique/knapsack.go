// 消化不良
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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()
	w, v, m := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		w[i], v[i], m[i] = getInt(), getInt(), getInt()
	}
	W := getInt()

	// DPテーブル（使い回す）
	dp := make([]int, W+1)

	// スライド最小値と同じ要領で最大値を求める
	deq := make([]int, W+1)  // デック（インデックス）
	deqv := make([]int, W+1) // デック（値）

	for i := 0; i < n; i++ {
		for a := 0; a < w[i]; a++ {
			s, t := 0, 0 // デックの先頭と末尾
			for j := 0; j*w[i]+a <= W; j++ {
				// デックの末尾にjを追加
				val := dp[j*w[i]+a] - j*v[i]
				for s < t && deqv[t-1] <= val {
					t--
				}
				deq[t] = j
				deqv[t] = val
				t++
				// デックの先頭を取り出す
				dp[j*w[i]+a] = deqv[s] + j*v[i]
				if deq[s] == j-m[i] {
					s++
				}
			}
		}
	}
	puts(dp[W])
}
