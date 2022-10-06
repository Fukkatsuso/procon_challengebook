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

// nまでの素数リストを返す
func PrimeNumbers(n int) []int {
	isPrime := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		isPrime[i] = true
	}
	isPrime[0], isPrime[1] = false, false

	prime := make([]int, 0)
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			prime = append(prime, i)
			for j := 2 * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	return prime
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	A, B, P := getInt(), getInt(), getInt()

	prime := PrimeNumbers(1000000)
	uf := NewUnionFind(B - A + 1)
	for _, p := range prime {
		if p < P {
			continue
		}
		start := (A + p - 1) / p * p // A以上の最小のpの倍数
		end := B / p * p             // B以下の最大のpの倍数
		for j := start; j <= end; j += p {
			uf.Unite(start-A, j-A)
		}
	}

	res := 0
	// 集合の個数を根の個数で数える
	for i := A; i <= B; i++ {
		if uf.Find(i-A) == i-A {
			res++
		}
	}
	puts(res)
}

type UnionFind struct {
	Par []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		Par: make([]int, n),
	}
	uf.Init(n)
	return uf
}

func (uf *UnionFind) Init(n int) {
	for i := 0; i < n; i++ {
		uf.Par[i] = -1
	}
}

func (uf *UnionFind) Find(x int) int {
	if uf.Par[x] < 0 {
		return x
	}
	uf.Par[x] = uf.Find(uf.Par[x])
	return uf.Par[x]
}

func (uf *UnionFind) Unite(x, y int) {
	x, y = uf.Find(x), uf.Find(y)
	if x == y {
		return
	}

	if uf.Par[x] > uf.Par[y] {
		x, y = y, x
	}
	uf.Par[x] += uf.Par[y]
	uf.Par[y] = x
}

func (uf *UnionFind) Same(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

// xの属する集合の要素数
func (uf *UnionFind) Size(x int) int {
	return -uf.Par[uf.Find(x)]
}
