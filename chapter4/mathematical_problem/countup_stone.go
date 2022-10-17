package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const (
	initialBufSize = 100000
	maxBufSize     = 10000000
	MOD            = 1000000007
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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := getInt(), getInt()

	primes := primeFactor(n)
	divs := divisor(n)

	res := 0
	for i := range divs {
		// オイラー関数のdivs[i]での値を求める
		euler := divs[i]
		for p := range primes {
			if divs[i]%p == 0 {
				euler = euler / p * (p - 1)
			}
		}

		res += euler * modPow(m, n/divs[i], MOD) % MOD
		res %= MOD
	}

	// 最後にnで割る
	// puts(res * modPow(n, MOD-2, MOD) % MOD)
	puts(res * modInv(n, MOD) % MOD)
}

func primeFactor(n int) map[int]int {
	m := map[int]int{}
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			m[i]++
			n /= i
		}
	}
	if n != 1 {
		m[n] = 1
	}
	return m
}

// 約数列挙
func divisor(n int) []int {
	div := make([]int, 0)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			div = append(div, i)
			if n/i != i {
				div = append(div, n/i)
			}
		}
	}
	sort.Ints(div)
	return div
}

func modPow(a, n, mod int) int {
	ret := 1
	for n > 0 {
		if n&1 == 1 {
			ret = (ret * a) % mod
		}
		a = (a * a) % mod
		n >>= 1
	}
	return ret
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
