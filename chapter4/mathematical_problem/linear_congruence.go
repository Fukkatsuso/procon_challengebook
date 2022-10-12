package main

// 連立線形合同式 Ax ≡ B (mod M) を解く
// 解全体は必ず x ≡ b (mod m) という形に書ける
// (b, m)を返す
func LinearCongruence(A, B, M []int) (int, int) {
	// 最初は条件がないので全ての整数を意味する x ≡ 0 (mod 1) を解としておく
	x, m := 0, 1
	for i := 0; i < len(A); i++ {
		a, b := A[i]*m, B[i]-A[i]*x
		d := gcd(M[i], a)
		if b%d != 0 {
			// 解なし
			return 0, -1
		}
		t := b / d * modInv(a/d, M[i]/d) % (M[i] / d)
		x += m * t
		m *= M[i] / d
	}
	return x % m, m
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func modInv(a, mod int) int {
	b := mod
	u, v := 1, 0
	for b > 0 {
		t := a / b
		a -= t * b
		u -= t * v
		a, b = b, a
		u, v = v, u
	}
	u %= mod
	if u < 0 {
		u += mod
	}
	return u
}
