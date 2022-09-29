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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, k := getInt(), getInt()
	price := make([][]int, n)
	for i := range price {
		price[i] = getInts(k)
	}

	V := n * 2
	g := NewGraph(V)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			upper := true
			for x := 0; x < k; x++ {
				upper = upper && price[i][x] > price[j][x]
			}
			if upper {
				g.AddEdge(i, n+j)
			}
		}
	}

	res := n - g.BipartiteMatching()
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
