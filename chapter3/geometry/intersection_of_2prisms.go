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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
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

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func minf(nums ...float64) float64 {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func maxf(nums ...float64) float64 {
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

	M, N := getInt(), getInt()
	X1, Y1 := make([]int, M), make([]int, M)
	for i := 0; i < M; i++ {
		X1[i], Y1[i] = getInt(), getInt()
	}
	X2, Z2 := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		X2[i], Z2[i] = getInt(), getInt()
	}

	// 区間の端点
	min1, max1 := min(X1...), max(X1...)
	min2, max2 := min(X2...), max(X2...)

	xs := make([]int, 0)
	for i := range X1 {
		xs = append(xs, X1[i])
	}
	for i := range X2 {
		xs = append(xs, X2[i])
	}
	sort.Ints(xs)

	f := func(x float64) float64 {
		return width(X1, Y1, M, x) * width(X2, Z2, N, x)
	}
	res := float64(0)
	for i := 0; i < len(xs)-1; i++ {
		a, b, c := float64(xs[i]), float64(xs[i+1]), float64(xs[i]+xs[i+1])/2
		// シンプソンの公式で積分
		if float64(min1) <= c && c <= float64(max1) &&
			float64(min2) <= c && c <= float64(max2) {
			res += (b - a) / 6 * (f(a) + 4*f(c) + f(b))
		}
	}
	putf("%.10f\n", res)
}

// 多角形をX=xでスライスしたときの幅
func width(X, Y []int, n int, x float64) float64 {
	lb, ub := float64(inf), float64(-inf)
	for i := 0; i < n; i++ {
		x1, y1 := float64(X[i]), float64(Y[i])
		x2, y2 := float64(X[(i+1)%n]), float64(Y[(i+1)%n])
		// i番目の辺と交点を持つか調べる
		if (x1-x)*(x2-x) <= 0 && x1 != x2 {
			// 交点の座標を計算
			y := y1 + (y2-y1)*(x-x1)/(x2-x1)
			lb = minf(lb, y)
			ub = maxf(ub, y)
		}
	}
	return maxf(0.0, ub-lb)
}
