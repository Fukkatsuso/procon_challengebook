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
	mod            = 10007
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

func multiply(A, B [][]int, mod int) [][]int {
	C := make([][]int, len(A))
	for i := range C {
		C[i] = make([]int, len(B[0]))
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B[0]); j++ {
			for k := 0; k < len(B); k++ {
				C[i][j] = (C[i][j] + A[i][k]*B[k][j]) % mod
			}
		}
	}
	return C
}

func modPow(A [][]int, n, mod int) [][]int {
	res := make([][]int, len(A))
	for i := range res {
		res[i] = make([]int, len(A))
		res[i][i] = 1
	}

	for n > 0 {
		if n&1 == 1 {
			res = multiply(res, A, mod)
		}
		A = multiply(A, A, mod)
		n >>= 1
	}
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()

	A := [][]int{
		{2, 1, 0},
		{2, 2, 2},
		{0, 1, 2},
	}
	A = modPow(A, n, mod)

	puts(A[0][0])
}
