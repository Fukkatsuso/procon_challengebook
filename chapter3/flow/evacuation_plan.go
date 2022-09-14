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
	E := make([][]int, N)
	for i := 0; i < N; i++ {
		E[i] = make([]int, M)
		for j := 0; j < M; j++ {
			E[i][j] = getInt()
		}
	}

	V := N + M + 1
	// 距離行列
	g := make([][]int, V)
	for i := 0; i < V; i++ {
		g[i] = make([]int, V)
		for j := 0; j < V; j++ {
			g[i][j] = INF
		}
	}
	for j := 0; j < M; j++ {
		sum := 0
		for i := 0; i < N; i++ {
			c := abs(X[i]-P[j]) + abs(Y[i]-Q[j]) + 1
			g[i][N+j] = c
			if E[i][j] > 0 {
				g[N+j][i] = -c
			}
			sum += E[i][j]
		}
		if sum > 0 {
			g[N+M][N+j] = 0
		}
		if sum < C[j] {
			g[N+j][N+M] = 0
		}
	}

	// ワーシャルフロイド法により負閉路検出を行う
	prev := make([][]int, V) // 最短路の直前の頂点
	for i := 0; i < V; i++ {
		prev[i] = make([]int, V)
		for j := 0; j < V; j++ {
			prev[i][j] = i
		}
	}
	for k := 0; k < V; k++ {
		for i := 0; i < V; i++ {
			for j := 0; j < V; j++ {
				if shorterPath := (g[i][j] > g[i][k]+g[k][j]); !shorterPath {
					continue
				}

				g[i][j] = g[i][k] + g[k][j]
				prev[i][j] = prev[k][j]

				if negativeCycle := (i == j && g[i][i] < 0); !negativeCycle {
					continue
				}

				// 負閉路が存在したら，閉路に沿ってフローを流す
				used := make([]bool, V) // ループ検出フラグ
				for v := i; !used[v]; v = prev[i][v] {
					used[v] = true
					if v != N+M && prev[i][v] != N+M {
						if v >= N {
							E[prev[i][v]][v-N]++
						} else {
							E[v][prev[i][v]-N]--
						}
					}
				}

				puts("SUBOPTIMAL")
				for x := 0; x < N; x++ {
					for y := 0; y < M; y++ {
						if y == M-1 {
							putf("%d\n", E[x][y])
						} else {
							putf("%d ", E[x][y])
						}
					}
				}
				return
			}
		}
	}

	// 最適な場合
	puts("OPTIMAL")
}
