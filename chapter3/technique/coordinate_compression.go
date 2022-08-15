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

type Point struct {
	x, y int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}

	w, h, n := getInt(), getInt(), getInt()
	x1 := getInts(n)
	x2 := getInts(n)
	y1 := getInts(n)
	y2 := getInts(n)

	w = compress(x1, x2, w)
	h = compress(y1, y2, h)

	// マス上に線があるか
	line := make([][]bool, h)
	for i := range line {
		line[i] = make([]bool, w)
	}
	for i := 0; i < n; i++ {
		for y := y1[i]; y <= y2[i]; y++ {
			for x := x1[i]; x <= x2[i]; x++ {
				line[y][x] = true
			}
		}
	}

	// 領域を数える
	res := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if line[y][x] {
				continue
			}

			res++

			// BFS
			q := make([]Point, 0)
			q = append(q, Point{x, y})
			for len(q) > 0 {
				p := q[0]
				q = q[1:]
				for d := 0; d < 4; d++ {
					tx, ty := p.x+dx[d], p.y+dy[d]
					if tx < 0 || w <= tx || ty < 0 || h <= ty {
						continue
					}
					if line[ty][tx] {
						continue
					}
					q = append(q, Point{tx, ty})
					line[ty][tx] = true
				}
			}
		}
	}

	puts(res)
}

// 座標圧縮
func compress(x1, x2 []int, w int) int {
	// 重複なく数える
	unique := func(x []int) map[int]bool {
		res := map[int]bool{}
		for i := range x {
			for d := -1; d <= 1; d++ {
				tx := x[i] + d
				if 1 <= tx && tx <= w {
					res[tx] = true
				}
			}
		}
		return res
	}

	// mapのキーを昇順ソートしたスライスを返す
	sorted := func(m map[int]bool) []int {
		res := make([]int, 0)
		for k := range m {
			res = append(res, k)
		}
		sort.Ints(res)
		return res
	}

	xs := sorted(unique(append(x1, x2...)))

	for i := 0; i < len(x1); i++ {
		x1[i] = sort.SearchInts(xs, x1[i])
		x2[i] = sort.SearchInts(xs, x2[i])
	}

	return len(xs)
}
