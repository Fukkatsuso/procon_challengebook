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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
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

	deq := make([]int, n) // デック
	s, t := 0, 0          // デックの先頭と末尾
	b := make([]int, n-k+1)
	for i := 0; i < n; i++ {
		// デックにiを追加
		for s < t && a[deq[t-1]] >= a[i] {
			t--
		}
		deq[t] = i
		t++

		if bi := i - k + 1; bi >= 0 {
			b[bi] = a[deq[s]]
			if deq[s] == bi {
				// デックの先頭を取り出す
				s++
			}
		}
	}

	for i := 0; i <= n-k; i++ {
		if i == n-k {
			putf("%d\n", b[i])
		} else {
			putf("%d ", b[i])
		}
	}
}
