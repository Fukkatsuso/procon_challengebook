// 強連結成分分解
// SCC: 強連結成分
package main

import "fmt"

type Graph struct {
	edge [][]int
	rev  [][]int
}

func NewGraph(v int) *Graph {
	g := Graph{
		edge: make([][]int, v),
		rev:  make([][]int, v),
	}

	for i := 0; i < v; i++ {
		g.edge[i] = make([]int, 0)
		g.rev[i] = make([]int, 0)
	}

	return &g
}

func (g *Graph) AddEdge(from, to int) {
	g.edge[from] = append(g.edge[from], to)
	g.rev[to] = append(g.rev[to], from)
}

// used: すでに調べたか
// vs: 帰りがけ順の並び
func (g *Graph) dfs(v int, used []bool, vs *[]int) {
	used[v] = true
	for _, to := range g.edge[v] {
		if !used[to] {
			g.dfs(to, used, vs)
		}
	}
	*vs = append(*vs, v)
}

func (g *Graph) rdfs(v, k int, used []bool, cmp []int) {
	used[v] = true
	cmp[v] = k
	for _, to := range g.rev[v] {
		if !used[to] {
			g.rdfs(to, k, used, cmp)
		}
	}
}

// 強連結成分(SCC)を求めるアルゴリズム
// 強連結成分の数, 各頂点が属する強連結成分のトポロジカル順序 を返す
func (g *Graph) SCC() (int, []int) {
	V := len(g.edge)
	used := make([]bool, V)
	vs := make([]int, 0)
	for v := 0; v < V; v++ {
		if !used[v] {
			g.dfs(v, used, &vs)
		}
	}

	for i := range used {
		used[i] = false
	}
	k := 0
	cmp := make([]int, V)
	for i := len(vs) - 1; i >= 0; i-- {
		if !used[vs[i]] {
			g.rdfs(vs[i], k, used, cmp)
			k++
		}
	}

	return k, cmp
}

// pp.285-286の例
func main() {
	g := NewGraph(12)
	edge := [][]int{
		{0, 1},
		{1, 2},
		{1, 3},
		{2, 3},
		{3, 4},
		{4, 2},
		{4, 5},
		{5, 6},
		{6, 7},
		{6, 8},
		{6, 9},
		{7, 5},
		{8, 9},
		{8, 11},
		{9, 10},
		{10, 9},
	}
	for i := range edge {
		from, to := edge[i][0], edge[i][1]
		g.AddEdge(from, to)
	}

	k, cmp := g.SCC()
	fmt.Println(k)
	fmt.Println(cmp)
}
