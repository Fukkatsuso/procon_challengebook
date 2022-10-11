package main

// オイラー関数の値を求める
// O(sqrt(n))
func EulerPhi(n int) int {
	res := n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			res = res / i * (i - 1)
			for n%i == 0 {
				n /= i
			}
		}
	}
	if n != 1 {
		res = res / n * (n - 1)
	}
	return res
}

// オイラー関数の値のテーブルを作成
// O(n)程度
func EulerPhiTable(n int) []int {
	euler := make([]int, n)
	for i := 0; i < n; i++ {
		euler[i] = i
	}
	for i := 2; i < n; i++ {
		if euler[i] == i {
			for j := i; j < n; j += i {
				euler[j] = euler[j] / i * (i - 1)
			}
		}
	}
	return euler
}
