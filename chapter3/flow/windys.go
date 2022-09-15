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

	N, M := getInt(), getInt()
	Z := make([][]int, N)
	for i := 0; i < N; i++ {
		Z[i] = getInts(M)
	}

	// 0~N: おもちゃ
	// N~2N-1: 0番の工場
	// 2N~3N-1: 1番の工場
	// ...
	// MN~(M+1)N-1: M-1番の工場
	s := (M + 1) * N
	t := s + 1
	V := t + 1
	g := NewGraph(V)
	// sからおもちゃiへの辺を張る
	for i := 0; i < N; i++ {
		g.AddEdge(s, i, 1, 0)
	}
	// 工場(j,k)からt，おもちゃiから工場(j,k)への辺を張る
	for j := 0; j < M; j++ {
		for k := 0; k < N; k++ {
			factory := (j+1)*N + k
			g.AddEdge(factory, t, 1, 0)
			for i := 0; i < N; i++ {
				cost := (k + 1) * Z[i][j]
				g.AddEdge(i, factory, 1, cost)
			}
		}
	}

	res := g.MinCostFlow(s, t, N)
	putf("%.6f\n", float64(res)/float64(N))
}

type Edge struct {
	to, cap, cost, rev int
}

type Graph [][]Edge

func NewGraph(v int) Graph {
	g := make([][]Edge, v)
	for i := range g {
		g[i] = make([]Edge, 0)
	}
	return g
}

func (g Graph) AddEdge(from, to, cap, cost int) {
	g[from] = append(g[from], Edge{
		to:   to,
		cap:  cap,
		cost: cost,
		rev:  len(g[to]),
	})
	g[to] = append(g[to], Edge{
		to:   from,
		cap:  0,
		cost: -cost,
		rev:  len(g[from]) - 1,
	})
}

// sからtへの流量fの最小費用流を求める
// 計算量はO(F|V||E|)
// 蟻本p.203にO(F|E|log|V|) or O(F|V|^2)の改良版アルゴリズムがある
func (g Graph) MinCostFlow(s, t, f int) int {
	const INF = 1 << 60

	prevv, preve := make([]int, len(g)), make([]int, len(g))

	res := 0
	for f > 0 {
		// ベルマンフォード法により，s-t間最短路を求める
		dist := make([]int, len(g))
		for i := range dist {
			dist[i] = INF
		}
		dist[s] = 0
		update := true
		for update {
			update = false
			for v := 0; v < len(g); v++ {
				if dist[v] == INF {
					continue
				}
				for i, e := range g[v] {
					if e.cap > 0 && dist[e.to] > dist[v]+e.cost {
						dist[e.to] = dist[v] + e.cost
						prevv[e.to], preve[e.to] = v, i
						update = true
					}
				}
			}
		}

		// これ以上流せない
		if dist[t] == INF {
			return -1
		}

		// s-t間最短路に沿って目一杯流す
		d := f
		for v := t; v != s; v = prevv[v] {
			d = min(d, g[prevv[v]][preve[v]].cap)
		}
		f -= d
		res += d * dist[t]
		for v := t; v != s; v = prevv[v] {
			e := &g[prevv[v]][preve[v]]
			e.cap -= d
			g[v][e.rev].cap += d
		}
	}
	return res
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
