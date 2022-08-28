package main

import "fmt"

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

// 蟻本p.188の例
func main() {
	g := NewGraph(5)
	g.AddEdge(0, 1, 10)
	g.AddEdge(0, 2, 2)
	g.AddEdge(1, 2, 6)
	g.AddEdge(1, 3, 6)
	g.AddEdge(2, 4, 5)
	g.AddEdge(3, 2, 3)
	g.AddEdge(3, 4, 8)

	flow := g.MaxFlow(0, 4)
	fmt.Println(flow)
}
