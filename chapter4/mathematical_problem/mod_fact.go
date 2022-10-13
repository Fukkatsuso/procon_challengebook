package main

import "fmt"

// n! = a p^e としたときの a mod p を求める
// O(log_p n)
// fact: n! mod p のテーブル．0<=n<p
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

func factTable(p int) []int {
	fact := make([]int, p)
	fact[0] = 1
	for i := 1; i < p; i++ {
		fact[i] = (fact[i-1] * i) % p
	}
	return fact
}

func main() {
	fmt.Println(factTable(7))
}
