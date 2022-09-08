package main

import "fmt"

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

// 蟻本p.199の例
func main() {
	g := NewGraph(5)
	g.AddEdge(0, 1, 10, 2)
	g.AddEdge(0, 2, 2, 4)
	g.AddEdge(1, 2, 6, 6)
	g.AddEdge(1, 3, 6, 2)
	g.AddEdge(2, 4, 5, 2)
	g.AddEdge(3, 2, 3, 3)
	g.AddEdge(3, 4, 8, 6)

	cost := g.MinCostFlow(0, 4, 9)
	fmt.Println(cost)
}
