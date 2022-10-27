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

	N, M := getInt(), getInt()
	V := N
	g := NewGraph(V)
	for i := 0; i < M; i++ {
		a, b := getInt(), getInt()
		g.AddEdge(a-1, b-1)
	}

	n, cmp := g.SCC()
	u, num := 0, 0 // u: トポロジカル順序で最後のSCCに含まれる頂点, num: 最後のSCCの頂点数=求める牛の数
	for v := 0; v < V; v++ {
		if cmp[v] == n-1 {
			u = v
			num++
		}
	}

	// すべての頂点から最後のSCCに到達可能か?
	used := make([]bool, V)
	g.rdfs(u, 0, used, cmp) // 頂点uから逆向きにDFSして，各頂点が最後のSCCに到達可能か調べる
	for v := 0; v < V; v++ {
		// 頂点vが到達不能
		if !used[v] {
			num = 0
			break
		}
	}

	puts(num)
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
