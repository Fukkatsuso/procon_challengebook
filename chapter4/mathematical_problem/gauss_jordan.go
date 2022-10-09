package main

import "fmt"

// Gauss-Jordanの消去法
// O(n^3), nは式の数

// Ax=bを解く
// Aは正方行列
// 解がないor一意でないならば，nilを返す
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

// p.256の例
func main() {
	A := [][]float64{
		{1, -2, 3},
		{4, -5, 6},
		{7, -8, 10},
	}
	b := []float64{
		6,
		12,
		21,
	}

	x := GaussJordan(A, b)
	if x == nil {
		fmt.Println("answer is nothing or ununique")
		return
	}
	for i := range x {
		fmt.Println(x[i])
	}
}
