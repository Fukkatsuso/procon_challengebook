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

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()
	h := getInts(n)

	stack := make([]int, n)

	L := make([]int, n) // iより左側で最も近い，h[i]より低い長方形のインデックス+1
	t := 0
	for i := 0; i < n; i++ {
		for t > 0 && h[stack[t-1]] > h[i] {
			t--
		}
		if t > 0 {
			L[i] = stack[t-1] + 1
		}
		stack[t] = i
		t++
	}

	R := make([]int, n) // iより右側で最も近い，h[i]より低い長方形のインデックス
	t = 0
	for i := n - 1; i >= 0; i-- {
		for t > 0 && h[stack[t-1]] >= h[i] {
			t--
		}
		if t > 0 {
			R[i] = stack[t-1]
		}
		stack[t] = i
		t++
	}

	res := 0
	for i := 0; i < n; i++ {
		res = max(res, h[i]*(R[i]-L[i]))
	}
	puts(res)
}
