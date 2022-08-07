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

func getInt2() (int, int) {
	return getInt(), getInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

type Item struct {
	w, v int
}

func (item Item) Eval(x float64) float64 {
	return float64(item.v) - x*float64(item.w)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, k := getInt2()
	items := make([]Item, n)
	for i := range items {
		w, v := getInt2()
		items[i] = Item{
			w: w,
			v: v,
		}
	}

	ok, ng := float64(0), float64(1<<60)
	for count := 0; count < 100; count++ {
		mid := (ok + ng) / 2

		sort.Slice(items, func(i, j int) bool {
			return items[i].Eval(mid) > items[j].Eval(mid)
		})

		sum := float64(0)
		for i := 0; i < k; i++ {
			sum += items[i].Eval(mid)
		}

		if sum >= 0 {
			ok = mid
		} else {
			ng = mid
		}
	}

	puts(ok)
}
