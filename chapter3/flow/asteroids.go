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

	n, k := getInt(), getInt()

	g := NewGraph(n * 2)
	for i := 0; i < k; i++ {
		r, c := getInt(), getInt()
		u, v := r-1, n+c-1
		g.AddEdge(u, v)
	}

	res := g.BipartiteMatching()
	puts(res)
}

type Graph [][]int

func NewGraph(v int) Graph {
	g := make(Graph, v)
	for i := range g {
		g[i] = make([]int, 0)
	}
	return g
}

func (g Graph) AddEdge(u, v int) {
	g[u] = append(g[u], v)
	g[v] = append(g[v], u)
}

// 増加パスを探す
func (g Graph) DFS(v int, match []int, used []bool) bool {
	used[v] = true
	for _, u := range g[v] {
		w := match[u]
		if w < 0 || (!used[w] && g.DFS(w, match, used)) {
			match[v] = u
			match[u] = v
			return true
		}
	}
	return false
}

// 二部グラフの最大マッチングを求める
func (g Graph) BipartiteMatching() int {
	res := 0
	match := make([]int, len(g))
	for i := range match {
		match[i] = -1
	}

	for v := 0; v < len(g); v++ {
		if match[v] < 0 {
			used := make([]bool, len(g))
			if g.DFS(v, match, used) {
				res++
			}
		}
	}
	return res
}
