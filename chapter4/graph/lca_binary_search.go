package main

import "math"

type Graph struct {
	edge [][]int
	root int
}

func NewGraph(n int, root int) *Graph {
	edge := make([][]int, n)
	for i := 0; i < n; i++ {
		edge[i] = make([]int, 0)
	}
	g := Graph{
		edge: edge,
		root: root,
	}
	return &g
}

type GraphLCA struct {
	graph  Graph
	parent [][]int
	depth  []int
}

func NewGraphLCA(g Graph) *GraphLCA {
	V := len(g.edge)
	maxLogV := int(math.Ceil(math.Log2(float64(V))))

	// parent[k]: 親を2^k回たどって到達する頂点（rootを通り過ぎる場合は-1）
	parent := make([][]int, maxLogV)
	for i := range parent {
		parent[i] = make([]int, V)
	}
	// rootからの深さ
	depth := make([]int, V)

	glca := GraphLCA{
		graph:  g,
		parent: parent,
		depth:  depth,
	}
	return &glca
}

func (g *GraphLCA) Init() {
	// parent[0]とdepthを初期化する
	g.dfs(g.graph.root, -1, 0)
	// parentを初期化する
	for k := 0; k+1 < len(g.parent); k++ {
		for v := range g.parent[k] {
			if g.parent[k][v] < 0 {
				g.parent[k+1][v] = -1
			} else {
				g.parent[k+1][v] = g.parent[k][g.parent[k][v]]
			}
		}
	}
}

func (g *GraphLCA) dfs(v, p, d int) {
	g.parent[0][v] = p
	g.depth[v] = d
	for _, to := range g.graph.edge[v] {
		if to != p {
			g.dfs(to, v, d+1)
		}
	}
}

// uとvのLCAを求める
// 二分探索を用いる手法
func (g *GraphLCA) LCA(u, v int) int {
	// uとvの深さが同じになるまで親をたどる
	if g.depth[u] > g.depth[v] {
		u, v = v, u
	}
	for k := range g.parent {
		if ((g.depth[v]-g.depth[u])>>k)&1 == 1 {
			v = g.parent[k][v]
		}
	}

	if u == v {
		return u
	}
	// 二分探索でLCAを求める
	// 2^k回ずつ親をたどっていく
	for k := len(g.parent) - 1; k >= 0; k-- {
		if g.parent[k][u] != g.parent[k][v] {
			u = g.parent[k][u]
			v = g.parent[k][v]
		}
	}
	return g.parent[0][u]
}
