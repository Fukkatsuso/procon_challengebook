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
	cows := gets()

	dir := make([]int, n)
	for i := 0; i < n; i++ {
		dir[i] = map[byte]int{
			'B': 1,
			'F': 0,
		}[cows[i]]
	}

	calc := func(k int) int {
		res := 0
		f := make([]int, n)
		fSum := 0
		for i := 0; i+k <= n; i++ {
			// 区間[i, i+K-1]
			// i番目の牛が後ろ向きなら反転させる
			if (dir[i]+fSum)%2 == 1 {
				res++
				f[i] = 1
			}
			// fSum更新
			fSum += f[i]
			if i-k+1 >= 0 {
				fSum -= f[i-k+1]
			}
		}

		// 反転できないn-k+1番目以降の牛は全て前向きか?
		for i := n - k + 1; i < n; i++ {
			if (dir[i]+fSum)%2 == 1 {
				return -1
			}
			if i-k+1 >= 0 {
				fSum -= f[i-k+1]
			}
		}

		return res
	}

	ansK, ansM := 1, n
	for k := 1; k <= n; k++ {
		m := calc(k)
		if m >= 0 && ansM >= m {
			ansK, ansM = k, m
		}
	}

	puts(ansK, ansM)
}
