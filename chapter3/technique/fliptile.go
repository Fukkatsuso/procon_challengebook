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
	inf            = 1 << 60
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

func newTile(m, n int) [][]int {
	tile := make([][]int, m)
	for i := range tile {
		tile[i] = make([]int, n)
	}
	return tile
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	m, n := getInt(), getInt()
	tile := make([][]int, m)
	for i := range tile {
		tile[i] = getInts(n)
	}

	res := inf           // 最小操作回数
	opt := newTile(m, n) // 最適解保存用

	for bit := 0; bit < (1 << n); bit++ {
		flip := newTile(m, n) // 作業用
		for j := 0; j < n; j++ {
			flip[0][j] = (bit >> (n - j - 1)) & 1
		}

		// 操作回数
		num := operation(m, n, tile, flip)
		if num >= 0 && num < res {
			res = num
			// copy
			for i := 0; i < m; i++ {
				for j := 0; j < n; j++ {
					opt[i][j] = flip[i][j]
				}
			}
		}
	}

	if res == inf {
		puts("IMPOSSIBLE")
		return
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == n-1 {
				putf("%d\n", opt[i][j])
			} else {
				putf("%d ", opt[i][j])
			}
		}
	}
}

func operation(m, n int, tile, flip [][]int) int {
	di := [5]int{-1, 0, 0, 0, 1}
	dj := [5]int{0, -1, 0, 1, 0}

	isBlack := func(i, j int) bool {
		c := tile[i][j]
		for d := 0; d < 5; d++ {
			i2, j2 := i+di[d], j+dj[d]
			if 0 <= i2 && i2 < m && 0 <= j2 && j2 < n {
				c += flip[i2][j2]
			}
		}
		return c%2 == 1
	}

	// 各マス，1行上が黒ならひっくり返す
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if isBlack(i-1, j) {
				flip[i][j] = 1
			}
		}
	}

	// 最終行が全て白か?
	for j := 0; j < n; j++ {
		if isBlack(m-1, j) {
			return -1
		}
	}

	// 反転回数をカウント
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res += flip[i][j]
		}
	}
	return res
}
