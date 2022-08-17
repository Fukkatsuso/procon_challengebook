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

	n := getInt()
	a := getInts(n)

	bit := NewBIT(n + 1)
	res := 0
	for i := 0; i < n; i++ {
		res += i - bit.sum(0, a[i]+1)
		bit.add(a[i], a[i]+1, 1)
	}

	puts(res)
}

// 1-indexed
type BIT [2][]int

func NewBIT(n int) *BIT {
	var b BIT
	for i := range b {
		b[i] = make([]int, n)
	}
	return &b
}

// [l,r)にxを加算
func (b *BIT) add(l, r, x int) {
	addSub := func(p, idx, x int) {
		for idx < len(b[p]) {
			b[p][idx] += x
			idx += idx & (-idx)
		}
	}
	addSub(0, l, -x*(l-1))
	addSub(0, r, x*(r-1))
	addSub(1, l, x)
	addSub(1, r, -x)
}

// [l,r)の和
func (b *BIT) sum(l, r int) int {
	sumSub := func(p, idx int) int {
		s := 0
		for idx > 0 {
			s += b[p][idx]
			idx -= idx & (-idx)
		}
		return s
	}
	return sumSub(0, r-1) + sumSub(1, r-1)*(r-1) - (sumSub(0, l-1) + sumSub(1, l-1)*(l-1))
}
