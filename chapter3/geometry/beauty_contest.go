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

	N := getInt()
	ps := make([]Point, N)
	for i := 0; i < N; i++ {
		x, y := getInt(), getInt()
		ps[i] = Point{
			x: x,
			y: y,
		}
	}

	qs := ConvexHull(ps)
	res := 0
	for i := range qs {
		for j := 0; j < i; j++ {
			res = max(res, Dist2(qs[i], qs[j]))
		}
	}
	puts(res)
}

type Point struct {
	x, y int
}

// 距離の2乗
func Dist2(p1, p2 Point) int {
	dx, dy := p1.x-p2.x, p1.y-p2.y
	return dx*dx + dy*dy
}

// p1-p2
func Sub(p1, p2 Point) Point {
	return Point{
		x: p1.x - p2.x,
		y: p1.y - p2.y,
	}
}

// 外積
func Det(p1, p2 Point) int {
	return p1.x*p2.y - p1.y*p2.x
}

// std::vector::resizeの実装
// https://cplusplus.com/reference/vector/vector/resize/
func Resize(ps []Point, n int) []Point {
	if n > len(ps) {
		for i := len(ps); i < n; i++ {
			ps = append(ps, ps[i-1])
		}
	}
	return ps[:n]
}

// 凸法を求める
func ConvexHull(ps []Point) []Point {
	sort.Slice(ps, func(i, j int) bool {
		if ps[i].x == ps[j].x {
			return ps[i].y < ps[j].y
		}
		return ps[i].x < ps[j].x
	})
	n := len(ps)
	k := 0                   // 凸包の頂点数
	qs := make([]Point, 2*n) // 構築中の凸包
	// 下側凸包の作成
	for i := 0; i < n; i++ {
		for k > 1 && Det(Sub(qs[k-1], qs[k-2]), Sub(ps[i], qs[k-1])) <= 0 {
			k--
		}
		qs[k] = ps[i]
		k++
	}
	// 上側凸包の作成
	for i, t := n-2, k; i >= 0; i-- {
		for k > t && Det(Sub(qs[k-1], qs[k-2]), Sub(ps[i], qs[k-1])) <= 0 {
			k--
		}
		qs[k] = ps[i]
		k++
	}
	return Resize(qs, k-1)
}
