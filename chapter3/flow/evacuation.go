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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	X, Y := getInt(), getInt()
	field := make([]string, X)
	for i := 0; i < X; i++ {
		field[i] = gets()
	}

	dX, dY := make([]int, 0), make([]int, 0) // ドアの座標
	pX, pY := make([]int, 0), make([]int, 0) // 人の座標
	dist := make([][][][]int, X)             // 最短距離
	for i := 0; i < X; i++ {
		dist[i] = make([][][]int, Y)
		for j := 0; j < Y; j++ {
			dist[i][j] = make([][]int, X)
			for k := 0; k < X; k++ {
				dist[i][j][k] = make([]int, Y)
				for l := 0; l < Y; l++ {
					dist[i][j][k][l] = -1
				}
			}
		}
	}

	// 各ドアからの最短距離を求める
	for x := 0; x < X; x++ {
		for y := 0; y < Y; y++ {
			if field[x][y] == 'D' {
				dX = append(dX, x)
				dY = append(dY, y)
				bfs(X, Y, x, y, field, dist[x][y])
			} else if field[x][y] == '.' {
				pX = append(pX, x)
				pY = append(pY, y)
			}
		}
	}

	// グラフを作成
	n := X * Y
	g := NewGraph(n * n)
	d, p := len(dX), len(pX)
	for i := 0; i < d; i++ { // 各ドアについて
		dx, dy := dX[i], dY[i]
		for j := 0; j < p; j++ { // 各人について
			px, py := pX[j], pY[j]
			if dist[dx][dy][px][py] >= 0 { // 人jがドアiに到達可能
				for k := dist[dx][dy][px][py]; k <= n; k++ { // 人jがドアiに時刻kの時点で到達する
					g.AddEdge((k-1)*d+i, n*d+j) // (時刻k,ドアi) <=> (人j)
				}
			}
		}
	}

	// 全員が脱出するための最小時間を計算
	if p == 0 {
		puts(0)
		return
	}
	num := 0
	match := make([]int, n*n)
	for i := range match {
		match[i] = -1
	}
	for v := 0; v < n*d; v++ {
		used := make([]bool, n*n)
		if g.DFS(v, match, used) {
			num++
			if num == p {
				puts(v/d + 1)
				return
			}
		}
	}

	// 脱出不能
	puts("impossible")
}

type Queue []int

func (q *Queue) empty() bool {
	return len(*q) == 0
}

func (q *Queue) push(x int) {
	*q = append(*q, x)
}

func (q *Queue) pop() int {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func bfs(X, Y, x, y int, field []string, d [][]int) {
	dx := [4]int{-1, 0, 0, 1}
	dy := [4]int{0, -1, 1, 0}

	qx, qy := make(Queue, 0), make(Queue, 0)
	d[x][y] = 0
	qx.push(x)
	qy.push(y)
	for !qx.empty() {
		x, y = qx.pop(), qy.pop()
		for k := 0; k < 4; k++ {
			x2, y2 := x+dx[k], y+dy[k]
			if 0 <= x2 && x2 < X &&
				0 <= y2 && y2 < Y &&
				field[x2][y2] == '.' &&
				d[x2][y2] < 0 {
				d[x2][y2] = d[x][y] + 1
				qx.push(x2)
				qy.push(y2)
			}
		}
	}
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
