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

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

type Item struct {
	w, v int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()
	w, v := getInts(n), getInts(n)
	W := getInt()

	// 前半を全列挙
	items := make([]Item, 1<<(n/2))
	for bit := 0; bit < (1 << (n / 2)); bit++ {
		wSum, vSum := 0, 0
		for i := 0; i < n/2; i++ {
			if (bit>>i)&1 == 1 {
				wSum += w[i]
				vSum += v[i]
			}
		}
		items[bit] = Item{
			w: wSum,
			v: vSum,
		}
	}

	// 無駄な要素を除く
	sort.Slice(items, func(i, j int) bool {
		if items[i].w == items[j].w {
			return items[i].v < items[j].v
		}
		return items[i].w < items[j].w
	})
	m := 1
	for i := 1; i < (1 << (n / 2)); i++ {
		if items[m-1].v < items[i].v {
			items[m] = items[i]
			m++
		}
	}

	// 後半を全列挙して解を求める
	res := 0
	for bit := 0; bit < (1 << (n - n/2)); bit++ {
		wSum, vSum := 0, 0
		for i := 0; i < (n - n/2); i++ {
			if (bit>>i)&1 == 1 {
				wSum += w[n/2+i]
				vSum += v[n/2+i]
			}
		}
		if wSum <= W {
			// 重さW-wSum以下となる最大のindex
			// <==> W-wSumを上回る最小のindexから1を引いたもの
			i := sort.Search(m, func(i int) bool {
				return items[i].w > W-wSum
			}) - 1
			res = max(res, vSum+items[i].v)
		}
	}

	puts(res)
}
