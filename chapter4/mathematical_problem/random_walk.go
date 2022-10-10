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
	dx = [4]int{-1, 1, 0, 0}
	dy = [4]int{0, 0, -1, 1}
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

	N, M := getInt(), getInt()
	grid := make([][]byte, N)
	for i := range grid {
		grid[i] = []byte(gets())
	}

	// canGoal[x][y]: (x,y)からゴールできるか
	canGoal := make([][]bool, N)
	for i := range canGoal {
		canGoal[i] = make([]bool, M)
	}
	dfs(N, M, canGoal, grid, N-1, M-1)

	// 行列の構築
	A := make([][]float64, N*M)
	for i := range A {
		A[i] = make([]float64, N*M)
	}
	b := make([]float64, N*M)
	for x := 0; x < N; x++ {
		for y := 0; y < M; y++ {
			if (x == N-1 && y == M-1) || !canGoal[x][y] {
				A[x*M+y][x*M+y] = 1
				continue
			}

			move := 0
			for k := 0; k < 4; k++ {
				nx, ny := x+dx[k], y+dy[k]
				if canMove(N, M, grid, nx, ny) {
					A[x*M+y][nx*M+ny] = -1
					move++
				}
			}
			b[x*M+y], A[x*M+y][x*M+y] = float64(move), float64(move)
		}
	}

	res := GaussJordan(A, b)
	putf("%.8f\n", res[0])
}

func dfs(N, M int, canGoal [][]bool, grid [][]byte, x, y int) {
	canGoal[x][y] = true
	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if canMove(N, M, grid, nx, ny) && !canGoal[nx][ny] {
			dfs(N, M, canGoal, grid, nx, ny)
		}
	}
}

func canMove(N, M int, grid [][]byte, x, y int) bool {
	return 0 <= x && x < N && 0 <= y && y < M && grid[x][y] == '.'
}

func GaussJordan(A [][]float64, b []float64) []float64 {
	n := len(A)
	B := make([][]float64, n)
	for i := 0; i < n; i++ {
		B[i] = make([]float64, n+1)
		for j := 0; j < n; j++ {
			B[i][j] = A[i][j]
		}
		// 行列Aの後ろにbを並べ同時に処理する
		B[i][n] = b[i]
	}

	const EPS = 1e-8
	for i := 0; i < n; i++ {
		// 注目している変数の係数の絶対値が大きい式をi番目に持ってくる
		pivot := i
		for j := i; j < n; j++ {
			if abs(B[i][j]) > abs(B[pivot][j]) {
				pivot = j
			}
		}
		B[i], B[pivot] = B[pivot], B[i]

		// 解がないor一意でない
		if abs(B[i][i]) < EPS {
			return nil
		}

		// 注目している変数の係数を1にする
		for j := i + 1; j <= n; j++ {
			B[i][j] /= B[i][i]
		}
		for j := 0; j < n; j++ {
			if i != j {
				// j番目の式からi番目の変数を消去
				for k := i + 1; k <= n; k++ {
					B[j][k] -= B[j][i] * B[i][k]
				}
			}
		}
	}

	x := make([]float64, n)
	// 後ろに並べたbが解になる
	for i := 0; i < n; i++ {
		x[i] = B[i][n]
	}
	return x
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
