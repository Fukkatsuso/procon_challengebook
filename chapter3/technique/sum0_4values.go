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

	n := getInt()
	A, B, C, D := getInts(n), getInts(n), getInts(n), getInts(n)

	ab, cd := combination(A, B), combination(C, D)

	ans := 0
	for k1, v1 := range ab {
		ans += v1 * cd[-k1]
	}

	puts(ans)
}

func combination(A, B []int) map[int]int {
	count := map[int]int{}
	for _, a := range A {
		for _, b := range B {
			count[a+b]++
		}
	}
	return count
}
