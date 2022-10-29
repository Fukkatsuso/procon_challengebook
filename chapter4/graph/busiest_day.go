package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	N := getInt()
	S, T, D := make([]int, N), make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		s, t, d := gets(), gets(), getInt()
		S[i] = toMinutes(s)
		T[i] = toMinutes(t)
		D[i] = d
	}

	// 0~N-1: x_i
	// N~2N-1: !x_i
	V := 2 * N
	g := NewGraph(V)
	// 結婚式iと結婚式j
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			// iの開始の式典 と jの開始の式典 が被る場合
			if min(S[i]+D[i], S[j]+D[j]) > max(S[i], S[j]) {
				// !x_i or !x_j <==> x_i => !x_j and x_j => !x_i
				g.AddEdge(i, N+j)
				g.AddEdge(j, N+i)
			}
			// iの開始の式典 と jの終了の式典 が被る場合
			if min(S[i]+D[i], T[j]) > max(S[i], T[j]-D[j]) {
				// !x_i or x_j <==> x_i => x_j and !x_j => !x_i
				g.AddEdge(i, j)
				g.AddEdge(N+j, N+i)
			}
			// iの終了の式典 と jの開始の式典 が被る場合
			if min(T[i], S[j]+D[j]) > max(T[i]-D[i], S[j]) {
				// x_i or !x_j <==> !x_i => !x_j and x_j => x_i
				g.AddEdge(N+i, N+j)
				g.AddEdge(j, i)
			}
			// iの終了の式典 と jの終了の式典 が被る場合
			if min(T[i], T[j]) > max(T[i]-D[i], T[j]-D[j]) {
				// x_i or x_j <==> !x_i => x_j and !x_j => x_i
				g.AddEdge(N+i, j)
				g.AddEdge(N+j, i)
			}
		}
	}
	_, cmp := g.SCC()

	// xと!xが異なる強連結成分に含まれるか判定
	for i := 0; i < N; i++ {
		if cmp[i] == cmp[N+i] {
			fmt.Println("NO")
			return
		}
	}

	// 充足可能な場合に解を復元
	// xを含むSCCのトポロジカル順序 が !xを含むSCCのトポロジカル順序 よりも後ろ なら，xには真を割り当てる
	fmt.Println("YES")
	for i := 0; i < N; i++ {
		if cmp[i] > cmp[N+i] {
			// x_iが真 <==> 結婚式iの最初に式典を行う
			puts(toClock(S[i]), toClock(S[i]+D[i]))
		} else {
			// x_iが偽 <==> 結婚式iの最後に式典を行う
			puts(toClock(T[i]-D[i]), toClock(T[i]))
		}
	}
}

func toMinutes(t string) int {
	s := strings.Split(t, ":")
	hour, _ := strconv.Atoi(s[0])
	minute, _ := strconv.Atoi(s[1])
	return hour*60 + minute
}

func toClock(minutes int) string {
	return fmt.Sprintf("%02d:%02d", minutes/60, minutes%60)
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
