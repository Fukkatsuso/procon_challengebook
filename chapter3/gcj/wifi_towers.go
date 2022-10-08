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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func square(x int) int {
	return x * x
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()
	x := getInts(n)
	y := getInts(n)
	r := getInts(n)
	s := getInts(n)

	S, T := n, n+1
	g := NewGraph(n + 2)

	ans := 0
	for i := 0; i < n; i++ {
		if s[i] < 0 {
			g.AddEdge(S, i, -s[i])
		} else {
			ans += s[i]
			g.AddEdge(i, T, s[i])
		}

		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if square(x[i]-x[j])+square(y[i]-y[j]) <= square(r[i]) {
				g.AddEdge(j, i, inf)
			}
		}
	}
	ans -= g.MaxFlow(S, T)

	puts(ans)
}

type Edge struct {
	to, cap, rev int
}

type Graph [][]Edge

func NewGraph(v int) Graph {
	g := make([][]Edge, v)
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

// sからの最短距離をBFSで計算する
// level: sからの距離
func (g Graph) BFS(s int, level []int) {
	for i := range level {
		level[i] = -1
	}
	que := make([]int, 0)
	level[s] = 0
	que = append(que, s)
	for len(que) != 0 {
		v := que[0]
		que = que[1:]
		for _, e := range g[v] {
			if e.cap > 0 && level[e.to] < 0 {
				level[e.to] = level[v] + 1
				que = append(que, e.to)
			}
		}
	}
}

// 増加パスをDFSで探す
// level: sからの距離
// iter: どこまで調べ終わったか
func (g Graph) DFS(v, t, f int, level, iter []int) int {
	if v == t {
		return f
	}
	for i := iter[v]; i < len(g[v]); i++ {
		e := &g[v][i]
		if e.cap > 0 && level[v] < level[e.to] {
			d := g.DFS(e.to, t, min(f, e.cap), level, iter)
			if d > 0 {
				e.cap -= d
				g[e.to][e.rev].cap += d
				return d
			}
		}
	}
	return 0
}

// sからtへの最大流を求める
func (g Graph) MaxFlow(s, t int) int {
	level := make([]int, len(g))
	const INF = 1 << 60

	flow := 0
	for {
		g.BFS(s, level)
		if level[t] < 0 {
			return flow
		}
		iter := make([]int, len(g))
		f := g.DFS(s, t, INF, level, iter)
		for f > 0 {
			flow += f
			f = g.DFS(s, t, INF, level, iter)
		}
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
