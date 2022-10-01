package main

import (
	"bufio"
	"fmt"
	"math"
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

func getInt2() (int, int) {
	return getInt(), getInt()
}

func getInt3() (int, int, int) {
	return getInt(), getInt(), getInt()
}

func getInt4() (int, int, int, int) {
	return getInt(), getInt(), getInt(), getInt()
}

func getInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = getInt()
	}
	return slice
}

func getFloat64() float64 {
	f, _ := strconv.ParseFloat(gets(), 64)
	return f
}

func getFloat64s(n int) []float64 {
	slice := make([]float64, n)
	for i := 0; i < n; i++ {
		slice[i] = getFloat64()
	}
	return slice
}

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	N := getInt()
	X, Y, R := make([]int, N), make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		X[i], Y[i], R[i] = getInt(), getInt(), getInt()
	}

	var low, high float64 = 0, 10000
	for i := 0; i < 100; i++ {
		mid := (low + high) / 2
		if ok(N, X, Y, R, mid) {
			high = mid
		} else {
			low = mid
		}
	}
	putf("%.6f\n", high)
}

func ok(N int, X, Y, R []int, r float64) bool {
	cand := make([]int, 0) // 1つの円で囲える集合
	cand = append(cand, 0)
	// pattern a
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			if float64(R[i]) >= r || float64(R[j]) >= r {
				continue
			}
			// 二円の交点を求める
			x1, y1, r1 := float64(X[i]), float64(Y[i]), r-float64(R[i])
			x2, y2, r2 := float64(X[j]), float64(Y[j]), r-float64(R[j])
			dx, dy := x2-x1, y2-y1
			a := dx*dx + dy*dy
			b := ((r1*r1-r2*r2)/a + 1) / 2
			d := r1*r1/a - b*b
			if d >= 0 {
				d = math.Sqrt(d)
				x3 := x1 + dx*b
				y3 := y1 + dy*b
				x4 := -dy * d
				y4 := dx * d
				// 誤差を考慮し，iとjは特別扱いする
				ij := (1 << uint(i)) | (1 << uint(j))
				cand = append(cand, cover(N, X, Y, R, x3-x4, y3-y4, r)|ij)
				cand = append(cand, cover(N, X, Y, R, x3+x4, y3+y4, r)|ij)
			}
		}
	}

	// pattern b
	for i := 0; i < N; i++ {
		if float64(R[i]) <= r {
			cand = append(cand, cover(N, X, Y, R, float64(X[i]), float64(Y[i]), r)|(1<<uint(i)))
		}
	}

	// 円の候補から2つ取り出し，全てを囲えているか調べる
	for i := range cand {
		for j := 0; j < i; j++ {
			if (cand[i] | cand[j]) == (1<<uint(N))-1 {
				return true
			}
		}
	}
	return false
}

// 中心(x,y)，半径rの円が囲う集合を求める
func cover(N int, X, Y, R []int, x, y, r float64) int {
	S := 0
	for i := 0; i < N; i++ {
		if float64(R[i]) <= r {
			dx, dy, dr := x-float64(X[i]), y-float64(Y[i]), r-float64(R[i])
			if dx*dx+dy*dy <= dr*dr {
				S |= 1 << uint(i)
			}
		}
	}
	return S
}
