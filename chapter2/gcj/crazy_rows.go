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
	// r[i]: i行目の一番右にある1のindex
	// 全ての行でr[i]<=iを満たすようにすればよい
	r := make([]int, n)
	for i := 0; i < n; i++ {
		line := gets()
		for j := 0; j < n; j++ {
			if line[j] == '1' {
				r[i] = j
			}
		}
	}

	ans := 0
	// i行目を確定させていく
	for i := 0; i < n; i++ {
		// r[j]<=i を満たすjのうち，最も小さいjをi行目に移動させる
		for j := i; j < n; j++ {
			if r[j] <= i {
				// k行目とk-1行目をswap
				for k := j; k-1 >= i; k-- {
					r[k], r[k-1] = r[k-1], r[k]
					ans++
				}
				break
			}
		}
	}

	puts(ans)
}
