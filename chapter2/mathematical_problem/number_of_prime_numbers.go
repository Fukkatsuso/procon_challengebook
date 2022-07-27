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

func Eratosthenes(n int) []bool {
	isPrime := make([]bool, n+1)

	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for i := 2; i <= n; i++ {
		if !isPrime[i] {
			continue
		}
		for j := i + i; j <= n; j += i {
			isPrime[j] = false
		}
	}

	return isPrime
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()

	isPrime := Eratosthenes(n)
	ans := 0
	for i := range isPrime {
		if isPrime[i] {
			ans++
		}
	}

	puts(ans)
}
