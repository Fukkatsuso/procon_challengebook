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
	INF            = 1 << 60
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

	N, K := getInt(), getInt()
	a, b, w := make([]int, N), make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		a[i], b[i], w[i] = getInt(), getInt(), getInt()
	}

	// 重複なし端点集合を求める
	points := map[int]bool{}
	for i := 0; i < N; i++ {
		points[a[i]] = true
		points[b[i]] = true
	}
	x := make([]int, 0)
	for k := range points {
		x = append(x, k)
	}
	sort.Ints(x)

	// グラフを作成
	m := len(x)
	s := m
	t := s + 1
	V := t + 1
	g := NewGraph(V)
	g.AddEdge(s, 0, K, 0)
	g.AddEdge(m-1, t, K, 0)
	for i := 0; i < m-1; i++ {
		g.AddEdge(i, i+1, INF, 0)
	}
	res := 0
	for i := 0; i < N; i++ {
		u := sort.SearchInts(x, a[i])
		v := sort.SearchInts(x, b[i])

		// uからvへ容量1，コスト-w[i]の辺を張る
		g.AddEdge(v, u, 1, w[i])
		g.AddEdge(s, v, 1, 0)
		g.AddEdge(u, t, 1, 0)
		res -= w[i]
	}

	res += g.MinCostFlow(s, t, K+N)
	puts(-res)
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
