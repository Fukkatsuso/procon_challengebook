package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	n := getInt()
	v1 := getInts(n)
	v2 := getInts(n)

	sort.Slice(v1, func(i, j int) bool {
		return v1[i] < v1[j]
	})
	sort.Slice(v2, func(i, j int) bool {
		return v2[i] > v2[j]
	})

	ans := 0
	for i := 0; i < n; i++ {
		ans += v1[i] * v2[i]
	}

	puts(ans)
}
