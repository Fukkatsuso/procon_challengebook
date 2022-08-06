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

func getInt2() (int, int) {
	return getInt(), getInt()
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

	n, m := getInt2()
	x := append([]int{-(1 << 60)}, getInts(n)...)
	sort.Ints(x)

	high, low := 100000, 1
	for high-low > 1 {
		mid := (high + low) / 2

		// 距離mid以上離した場合，何頭の牛が牛舎に入るか
		k := 0
		// 直前に牛を入れた牛舎の番号
		prev := 0
		for i := 1; i <= n; i++ {
			if x[i]-x[prev] >= mid {
				k++
				prev = i
			}
		}

		if k >= m {
			low = mid
		} else {
			high = mid
		}
	}

	puts(low)
}
