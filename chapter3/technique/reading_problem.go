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

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	p := getInt()
	a := getInts(p)

	// 出てくる事柄の種類数を計算
	count := map[int]int{}
	for i := range a {
		count[a[i]]++
	}
	allKind := len(count)

	count = map[int]int{}
	appear := 0
	ans := p
	for l, r := 0, 0; l < p; l++ {
		for ; r < p && appear < allKind; r++ {
			if count[a[r]] == 0 {
				appear++
			}
			count[a[r]]++
		}

		if appear == allKind {
			ans = min(ans, r-l)
		}

		count[a[l]]--
		if count[a[l]] == 0 {
			appear--
		}
	}

	puts(ans)
}
