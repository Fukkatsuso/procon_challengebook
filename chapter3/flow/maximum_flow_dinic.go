package main

// Dinic法
// 最短の増加パスを探してそこにフローを流す

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
