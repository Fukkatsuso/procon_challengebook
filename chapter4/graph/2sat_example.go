package main

import "fmt"

func main() {
	V := 6
	g := NewGraph(V)

	v := map[string]int{
		"a":  0,
		"b":  1,
		"c":  2,
		"!a": 3,
		"!b": 4,
		"!c": 5,
	}
	// aまたは!b は !a=>!bかつb=>a
	g.AddEdge(v["!a"], v["!b"])
	g.AddEdge(v["b"], v["a"])
	// bまたはc は !b=>cかつ!c=>b
	g.AddEdge(v["!b"], v["c"])
	g.AddEdge(v["!c"], v["b"])
	// !cまたは!a は c=>!aかつa=>!c
	g.AddEdge(v["c"], v["!a"])
	g.AddEdge(v["a"], v["!c"])

	_, cmp := g.SCC()

	// xと!xが異なる強連結成分に含まれるか判定
	for i := 0; i < 3; i++ {
		if cmp[i] == cmp[3+i] {
			fmt.Println("NO")
			return
		}
	}

	// 充足可能な場合に解を復元
	// xを含むSCCのトポロジカル順序 が !xを含むSCCのトポロジカル順序 よりも後ろ なら，xには真を割り当てる
	fmt.Println("YES")
	for i := 0; i < 3; i++ {
		if cmp[i] > cmp[3+i] {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	}
}

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
