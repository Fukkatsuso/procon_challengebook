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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const ST_SIZE = (1 << 15) - 1

	n, c := getInt(), getInt()
	L := getInts(n)
	S, A := getInts(c), getInts(c)

	// セグメント木のデータ
	seg := Seg{
		vx:  make([]float64, ST_SIZE),
		vy:  make([]float64, ST_SIZE),
		ang: make([]float64, ST_SIZE),
	}

	// 角度の変化を調べるため，現在の角度を保存しておく
	prev := make([]float64, n)

	seg.init(0, 0, n, L)
	for i := 1; i < n; i++ {
		prev[i] = math.Pi
	}

	// 各クエリを処理
	for i := 0; i < c; i++ {
		s := S[i]
		a := float64(A[i]) / 360.0 * 2.0 * math.Pi

		seg.change(s, a-prev[s], 0, 0, n)
		prev[s] = a

		putf("%.2f %.2f\n", seg.vx[0], seg.vy[0])
	}
}

type Seg struct {
	vx, vy, ang []float64
}

func (seg *Seg) init(k, l, r int, L []int) {
	seg.ang[k], seg.vx[k] = 0.0, 0.0

	if r-l == 1 {
		seg.vy[k] = float64(L[l])
	} else {
		chl, chr := k*2+1, k*2+2
		mid := (l + r) / 2
		seg.init(chl, l, mid, L)
		seg.init(chr, mid, r, L)
		seg.vy[k] = seg.vy[chl] + seg.vy[chr]
	}
}

// 場所sの角度をaだけ変更
func (seg *Seg) change(s int, a float64, v, l, r int) {
	if s <= l || r <= s {
		return
	}

	chl, chr := v*2+1, v*2+2
	mid := (l + r) / 2
	seg.change(s, a, chl, l, mid)
	seg.change(s, a, chr, mid, r)
	if s <= mid {
		seg.ang[v] += a
	}

	sin, cos := math.Sin(seg.ang[v]), math.Cos(seg.ang[v])
	seg.vx[v] = seg.vx[chl] + cos*seg.vx[chr] - sin*seg.vy[chr]
	seg.vy[v] = seg.vy[chl] + sin*seg.vx[chr] + cos*seg.vy[chr]
}
