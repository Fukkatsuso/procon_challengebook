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
	p := getInts(n)

	if n%2 == 1 {
		// 盤面の左端を駒とみなす
		p = append(p, 0)
		n++
	}
	sort.Ints(p)

	x := 0
	for i := 0; i+1 < n; i += 2 {
		x ^= (p[i+1] - p[i] - 1)
	}

	if x == 0 {
		puts("Bob will win")
	} else {
		puts("Georgia will win")
	}
}
