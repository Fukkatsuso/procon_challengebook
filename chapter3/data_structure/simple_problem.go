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

	n, q := getInt(), getInt()
	a := getInts(n)

	bit := NewBIT(n + 1)
	for i := 1; i <= n; i++ {
		bit.add(i, i+1, a[i-1])
	}

	for i := 0; i < q; i++ {
		op := gets()
		if op == "Q" {
			l, r := getInt(), getInt()
			puts(bit.sum(l, r+1))
		} else {
			l, r, x := getInt(), getInt(), getInt()
			bit.add(l, r+1, x)
		}
	}
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

// http://poj.org/problem?id=3468
// Sample Input
// 10 5
// 1 2 3 4 5 6 7 8 9 10
// Q 4 4
// Q 1 10
// Q 2 4
// C 3 6 3
// Q 2 4

// Sample Output
// 4
// 55
// 9
// 15
