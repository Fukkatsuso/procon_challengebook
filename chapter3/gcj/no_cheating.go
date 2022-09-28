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

	dx := []int{-1, 1, -1, 1}
	dy := []int{0, 0, -1, -1}

	M, N := getInt(), getInt()
	seat := make([]string, M)
	for i := range seat {
		seat[i] = gets()
	}

	V := N * M
	g := NewGraph(V)
	canSit := 0
	for y := 0; y < M; y++ {
		for x := 0; x < N; x++ {
			if seat[y][x] == 'x' {
				continue
			}
			canSit++
			for k := 0; k < 4; k++ {
				x2, y2 := x+dx[k], y+dy[k]
				if 0 <= x2 && x2 < N &&
					0 <= y2 && y2 < M &&
					seat[y2][x2] == '.' {
					g.AddEdge(y*M+x, y2*M+x2)
				}
			}
		}
	}
	puts(canSit - g.BipartiteMatching())
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
