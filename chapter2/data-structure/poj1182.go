// x
package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	T := make([]int, k)
	X := make([]int, k)
	Y := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&T[i], &X[i], &Y[i])
	}

	par := make([]int, 3*n)
	rank := make([]int, 3*n)

	UFInit(par, rank, 3*n)

	ans := 0
	for i := 0; i < k; i++ {
		t := T[i]
		x := X[i] - 1
		y := Y[i] - 1

		// 正しくない番号
		if x < 0 || n <= x || y < 0 || n <= y {
			ans++
			continue
		}

		if t == 1 {
			// 「x,yが同じ種類」という情報
			if UFSame(par, x, y+n) || UFSame(par, x, y+2*n) {
				ans++
			} else {
				UFUnite(par, rank, x, y)
				UFUnite(par, rank, x+n, y+n)
				UFUnite(par, rank, x+2*n, y+2*n)
			}
		} else {
			// 「xがyを食べる」という情報
			if UFSame(par, x, y) || UFSame(par, x, y+2*n) {
				ans++
			} else {
				UFUnite(par, rank, x, y+n)
				UFUnite(par, rank, x+n, y+2*n)
				UFUnite(par, rank, x+2*n, y)
			}
		}
	}

	fmt.Println(ans)
}

func UFInit(par, rank []int, n int) {
	for i := 0; i < n; i++ {
		par[i] = i
		rank[i] = 0
	}
}

func UFFind(par []int, x int) int {
	for par[x] != x {
		x, par[x] = par[x], par[par[x]]
	}
	return x
}

func UFUnite(par, rank []int, x, y int) {
	x = UFFind(par, x)
	y = UFFind(par, y)
	if x == y {
		return
	}

	if rank[x] < rank[y] {
		par[x] = y
	} else {
		par[y] = x
		if rank[x] == rank[y] {
			rank[x]++
		}
	}
}

func UFSame(par []int, x, y int) bool {
	return UFFind(par, x) == UFFind(par, y)
}
