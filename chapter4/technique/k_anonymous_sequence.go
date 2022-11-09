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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, k := getInt(), getInt()
	a := getInts(n)

	dp := make([]int, n+1) // DPテーブル
	S := make([]int, n+1)  // aの累積和
	deq := make([]int, n)  // デック

	// 直線f_jのxにおける値
	f := func(j, x int) int {
		return -a[j]*x + dp[j] - S[j] + a[j]*j
	}

	// f2が最小値をとる可能性があるか判定
	check := func(f1, f2, f3 int) bool {
		a1, b1 := -a[f1], dp[f1]-S[f1]+a[f1]*f1
		a2, b2 := -a[f2], dp[f2]-S[f2]+a[f2]*f2
		a3, b3 := -a[f3], dp[f3]-S[f3]+a[f3]*f3
		return (a2-a1)*(b3-b2) >= (b2-b1)*(a3-a2)
	}

	// 累積和の計算
	for i := 0; i < n; i++ {
		S[i+1] = S[i] + a[i]
	}

	// デックの初期化
	s, t := 0, 1
	deq[0] = 0

	dp[0] = 0

	for i := k; i <= n; i++ {
		if i-k >= k {
			// 末尾から最小値をとる可能性がなくなったものを取り除く
			for s+1 < t && check(deq[t-2], deq[t-1], i-k) {
				t--
			}

			// デックにi-kを追加
			deq[t] = i - k
			t++
		}

		// 先頭が最小値でなければ取り除く
		for s+1 < t && f(deq[s], i) >= f(deq[s+1], i) {
			s++
		}

		dp[i] = S[i] + f(deq[s], i)
	}

	puts(dp[n])
}
