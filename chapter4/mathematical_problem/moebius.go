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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const MOD = 10009

	n := getInt()

	res := 0
	mu := Moebius(n)
	for k, v := range mu {
		res += v * modPow(26, n/k, MOD) // μ(d) * 26^(n/d)
		res = (res%MOD + MOD) % MOD
	}
	puts(res)
}

// nの約数におけるメビウス関数の値のmapを返す
// O(sqrt(n))
func Moebius(n int) map[int]int {
	primes := make([]int, 0)

	// nの素因数を列挙する
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			primes = append(primes, i)
			for n%i == 0 {
				n /= i
			}
		}
	}
	if n != 1 {
		primes = append(primes, n)
	}

	res := map[int]int{}
	m := len(primes)
	for i := 0; i < (1 << m); i++ {
		mu, d := 1, 1
		for j := 0; j < m; j++ {
			if (i>>j)&1 == 1 {
				mu *= -1
				d *= primes[j]
			}
		}
		res[d] = mu
	}
	return res
}
