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

	n, k := getInt(), getInt()
	a := getInts(k)
	x := getInts(n)

	g := grundy(x, a)

	// 勝敗を判定
	t := 0
	for i := range x {
		t ^= g[x[i]]
	}

	if t != 0 {
		puts("Alice")
	} else {
		puts("Bob")
	}
}

func grundy(x, a []int) []int {
	maxX := max(x...)
	dp := make([]int, maxX+1)
	// 0枚で自分に回ってきたら負け
	dp[0] = 0
	// grundy数を計算
	for j := 1; j <= maxX; j++ {
		// s: 状態jから到達可能な状態j-a[i] の集合
		s := map[int]int{}
		for i := range a {
			if a[i] <= j {
				s[dp[j-a[i]]]++
			}
		}

		// g: sに含まれない最小の非負の整数
		g := 0
		for s[g] != 0 {
			g++
		}
		dp[j] = g
	}

	return dp
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
