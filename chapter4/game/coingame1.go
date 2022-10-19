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

	x, k := getInt(), getInt()
	a := getInts(k)

	const (
		WIN  = true
		LOSE = false
	)

	win := make([]bool, x+1)
	// 0枚で自分に返ってきたら負け
	win[0] = LOSE
	for i := 1; i <= x; i++ {
		win[i] = LOSE
		// 相手を負けにできれば自分の勝ち
		for j := range a {
			win[i] = win[i] || (i >= a[j] && !win[i-a[j]])
		}
	}

	if win[x] {
		puts("Alice")
	} else {
		puts("Bob")
	}
}
