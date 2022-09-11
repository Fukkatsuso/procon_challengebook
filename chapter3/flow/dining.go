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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	N, F, D := getInt(), getInt(), getInt()

	likeF, likeD := make([][]bool, N), make([][]bool, N)
	for i := range likeF {
		likeF[i] = make([]bool, F)
		likeD[i] = make([]bool, D)
	}
	for i := 0; i < N; i++ {
		fn, dn := getInt(), getInt()
		for j := 0; j < fn; j++ {
			f := getInt() - 1
			likeF[i][f] = true
		}
		for j := 0; j < dn; j++ {
			d := getInt() - 1
			likeD[i][d] = true
		}
	}

	// 0~N-1: 食べ物側の牛
	cowF := func(i int) int {
		return i
	}
	// N~2N-1: 飲み物側の牛
	cowD := func(i int) int {
		return N + i
	}
	// 2N~2N+F-1: 食べ物
	food := func(i int) int {
		return 2*N + i
	}
	// 2N+F~2N+F+D-1: 飲み物
	drink := func(i int) int {
		return 2*N + F + i
	}
	s := 2*N + F + D
	t := s + 1
	g := NewGraph(t + 1)

	// sと食べ物を結ぶ
	for i := 0; i < F; i++ {
		g.AddEdge(s, food(i), 1)
	}

	// 飲み物とtを結ぶ
	for i := 0; i < D; i++ {
		g.AddEdge(drink(i), t, 1)
	}

	for i := 0; i < N; i++ {
		// 食べ物側の牛と飲み物側の牛を結ぶ
		g.AddEdge(cowF(i), cowD(i), 1)

		for j := 0; j < F; j++ {
			if likeF[i][j] {
				// 食べ物と牛を結ぶ
				g.AddEdge(food(j), cowF(i), 1)
			}
		}
		for j := 0; j < D; j++ {
			if likeD[i][j] {
				// 牛と飲み物を結ぶ
				g.AddEdge(cowD(i), drink(j), 1)
			}
		}
	}

	res := g.MaxFlow(s, t)
	puts(res)
}

type Edge struct {
	to, cap, rev int
}

type Graph [][]Edge

func NewGraph(n int) Graph {
	g := make([][]Edge, n)
	for i := range g {
		g[i] = make([]Edge, 0)
	}
	return g
}

func (g Graph) AddEdge(from, to, cap int) {
	g[from] = append(g[from], Edge{
		to:  to,
		cap: cap,
		rev: len(g[to]),
	})
	g[to] = append(g[to], Edge{
		to:  from,
		cap: 0,
		rev: len(g[from]) - 1,
	})
}

func (g Graph) DFS(v, t, f int, used []bool) int {
	if v == t {
		return f
	}
	used[v] = true
	for i := range g[v] {
		e := &g[v][i]
		if !used[e.to] && e.cap > 0 {
			d := g.DFS(e.to, t, min(f, e.cap), used)
			if d > 0 {
				e.cap -= d
				g[e.to][e.rev].cap += d
				return d
			}
		}
	}
	return 0
}

func (g Graph) MaxFlow(s, t int) int {
	flow := 0
	for {
		used := make([]bool, len(g))
		f := g.DFS(s, t, 1<<60, used)
		if f == 0 {
			return flow
		}
		flow += f
	}
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
