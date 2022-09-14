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
	INF            = 1 << 60
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	N, M := getInt(), getInt()
	X, Y, B := make([]int, N), make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		X[i], Y[i], B[i] = getInt(), getInt(), getInt()
	}
	P, Q, C := make([]int, M), make([]int, M), make([]int, M)
	for i := 0; i < M; i++ {
		P[i], Q[i], C[i] = getInt(), getInt(), getInt()
	}

	// マッチンググラフを作成
	// 0~N-1: ビル
	// N~N+M-1: シェルター
	s := N + M
	t := s + 1
	V := t + 1
	g := NewGraph(V)
	cost := 0 // 避難計画のコストの総和
	F := 0    // 人の総数
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			e := getInt()
			c := abs(X[i]-P[j]) + abs(Y[i]-Q[j]) + 1
			g.AddEdge(i, N+j, INF, c)
			cost += c * e
		}
	}
	for i := 0; i < N; i++ {
		g.AddEdge(s, i, B[i], 0)
		F += B[i]
	}
	for i := 0; i < M; i++ {
		g.AddEdge(N+i, t, C[i], 0)
	}

	// 最適な場合
	if g.MinCostFlow(s, t, F) == cost {
		puts("OPTIMAL")
		return
	}

	// 最適でない場合
	puts("SUBOPTIMAL")
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			// ビルi->シェルターj の流量は，逆辺のcapを見ればわかる
			if j == M-1 {
				putf("%d\n", g[N+j][i].cap)
			} else {
				putf("%d ", g[N+j][i].cap)
			}
		}
	}
}

type Edge struct {
	to, cap, cost, rev int
}

type Graph [][]Edge

func NewGraph(v int) Graph {
	g := make([][]Edge, v)
	for i := range g {
		g[i] = make([]Edge, 0)
	}
	return g
}

func (g Graph) AddEdge(from, to, cap, cost int) {
	g[from] = append(g[from], Edge{
		to:   to,
		cap:  cap,
		cost: cost,
		rev:  len(g[to]),
	})
	g[to] = append(g[to], Edge{
		to:   from,
		cap:  0,
		cost: -cost,
		rev:  len(g[from]) - 1,
	})
}

// sからtへの流量fの最小費用流を求める
// 計算量はO(F|V||E|)
// 蟻本p.203にO(F|E|log|V|) or O(F|V|^2)の改良版アルゴリズムがある
func (g Graph) MinCostFlow(s, t, f int) int {
	const INF = 1 << 60

	prevv, preve := make([]int, len(g)), make([]int, len(g))

	res := 0
	for f > 0 {
		// ベルマンフォード法により，s-t間最短路を求める
		dist := make([]int, len(g))
		for i := range dist {
			dist[i] = INF
		}
		dist[s] = 0
		update := true
		for update {
			update = false
			for v := 0; v < len(g); v++ {
				if dist[v] == INF {
					continue
				}
				for i, e := range g[v] {
					if e.cap > 0 && dist[e.to] > dist[v]+e.cost {
						dist[e.to] = dist[v] + e.cost
						prevv[e.to], preve[e.to] = v, i
						update = true
					}
				}
			}
		}

		// これ以上流せない
		if dist[t] == INF {
			return -1
		}

		// s-t間最短路に沿って目一杯流す
		d := f
		for v := t; v != s; v = prevv[v] {
			d = min(d, g[prevv[v]][preve[v]].cap)
		}
		f -= d
		res += d * dist[t]
		for v := t; v != s; v = prevv[v] {
			e := &g[prevv[v]][preve[v]]
			e.cap -= d
			g[v][e.rev].cap += d
		}
	}
	return res
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
