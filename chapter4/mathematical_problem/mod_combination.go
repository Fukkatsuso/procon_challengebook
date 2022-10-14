package main

// nCk mod p
// O(log_p n)
// fact: n! mod p のテーブル．実装はmod_fact.goを参照
func ModComb(n, k, p int, fact []int) int {
	if n < 0 || k < 0 || n < k {
		return 0
	}
	var e1, e2, e3 int
	a1, a2, a3 := ModFact(n, p, &e1, fact), ModFact(k, p, &e2, fact), ModFact(n-k, p, &e3, fact)
	if e1 > e2+e3 {
		return 0
	}
	return a1 * ModInv(a2*a3%p, p) % p
}

func ModFact(n, p int, e *int, fact []int) int {
	*e = 0
	if n == 0 {
		return 1
	}

	// pの倍数の部分を計算
	res := ModFact(n/p, p, e, fact)
	*e += n / p

	// (p-1)! ≡ -1 なので (p-1)!^(n/p) は n/p の偶奇だけで計算可能
	if (n/p)%2 != 0 {
		return res * (p - fact[n%p]) % p
	}
	return res * fact[n%p] % p
}

func ModInv(a, mod int) int {
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
