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
	BUCKET_SIZE    = 1000
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

	n, m := getInt(), getInt()
	a := getInts(n)

	aSorted := make([]int, n)                              // aをソートしたもの
	bucket := make([][]int, (n+BUCKET_SIZE-1)/BUCKET_SIZE) // 各バケットをソートしたもの
	for i := 0; i < n; i++ {
		aSorted[i] = a[i]
		if bucket[i/BUCKET_SIZE] == nil {
			bucket[i/BUCKET_SIZE] = make([]int, 0)
		}
		bucket[i/BUCKET_SIZE] = append(bucket[i/BUCKET_SIZE], a[i])
	}
	sort.Ints(aSorted)
	for i := range bucket {
		sort.Ints(bucket[i])
	}

	for i := 0; i < m; i++ {
		// [l, r)のk番目の数を求める
		l, r, k := getInt()-1, getInt(), getInt()

		low, high := -1, n-1
		for high-low > 1 {
			mid := (low + high) / 2
			// x以下の数をカウントしていく
			x := aSorted[mid]
			tl, tr, c := l, r, 0

			// バケットをはみ出す部分をカウント
			for tl < tr && tl%BUCKET_SIZE != 0 {
				if a[tl] <= x {
					c++
				}
				tl++
			}
			for tl < tr && tr%BUCKET_SIZE != 0 {
				tr--
				if a[tr] <= x {
					c++
				}
			}

			// バケット内をカウント
			for tl < tr {
				b := tl / BUCKET_SIZE
				c += sort.Search(len(bucket[b]), func(i int) bool {
					return bucket[b][i] > x
				})
				tl += BUCKET_SIZE
			}

			if c >= k {
				high = mid
			} else {
				low = mid
			}
		}

		puts(aSorted[high])
	}
}
