package main

import (
	"bufio"
	"container/heap"
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

func getInt2() (int, int) {
	return getInt(), getInt()
}

func getInt3() (int, int, int) {
	return getInt(), getInt(), getInt()
}

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

// ある頂点vへの2番目の最短距離は，
// 他の頂点uへの最短路にu->vをつなげたものか，
// uへの2番めの最短路にu->vをつなげたもの
func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, r := getInt2()
	g := NewGraph(n)
	for i := 0; i < r; i++ {
		from, to, cost := getInt3()
		g.AddEdge(from-1, to-1, cost)
		g.AddEdge(to-1, from-1, cost)
	}

	g.DijkstraSearch(0)

	puts(g.Cost2[n-1])
}

type Edge struct {
	to, cost int
}

// 辺のリストで表現されるグラフ
type Graph struct {
	Edges [][]Edge // Edge[i][j]: 頂点iのj番目の辺
	Cost  []int    // 始点からの最小コスト
	Cost2 []int    // 2番目の最小コスト
}

// 頂点数nのリスト型グラフを返す
func NewGraph(n int) *Graph {
	g := &Graph{
		Edges: make([][]Edge, n),
		Cost:  make([]int, n),
		Cost2: make([]int, n),
	}
	for i := 0; i < n; i++ {
		g.Edges[i] = make([]Edge, 0)
		g.Cost[i] = 1 << 60
		g.Cost2[i] = 1 << 60
	}
	return g
}

// 辺の追加
func (g *Graph) AddEdge(from, to, cost int) {
	g.Edges[from] = append(g.Edges[from], Edge{to, cost})
}

// value: 頂点の番号
// priority: 始点からのコスト
type Elem struct {
	value    interface{}
	priority int
	index    int
}

type PriorityQueue []*Elem

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	elem := x.(*Elem)
	elem.index = n
	*pq = append(*pq, elem)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(*pq)
	elem := old[n-1]
	old[n-1] = nil
	elem.index = -1
	*pq = old[0 : n-1]
	return elem
}

// 0-indexed
func (g *Graph) DijkstraSearch(origin int) {
	// init
	g.Cost[origin] = 0
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Elem{
		value:    origin,
		priority: 0,
	})

	// seaech
	for pq.Len() > 0 {
		v := heap.Pop(&pq).(*Elem)
		from := v.value.(int)
		if g.Cost2[from] < v.priority {
			continue
		}
		for _, edge := range g.Edges[from] {
			to := edge.to
			// 隣接頂点の最小コストを更新
			nextCost := v.priority + edge.cost
			if nextCost < g.Cost[to] {
				g.Cost[to] = nextCost
				heap.Push(&pq, &Elem{
					value:    to,
					priority: nextCost,
				})
			}
			// 2番目の最小コストを更新
			if nextCost < g.Cost2[to] && nextCost > g.Cost[to] {
				g.Cost2[to] = nextCost
				heap.Push(&pq, &Elem{
					value:    to,
					priority: nextCost,
				})
			}
		}
	}
}
