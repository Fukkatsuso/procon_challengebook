package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

func getInt4() (int, int, int, int) {
	return getInt(), getInt(), getInt(), getInt()
}

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	N, H, R, T := getInt4()

	y := make([]float64, N)
	for i := 0; i < N; i++ {
		y[i] = height(float64(H), T-i)
	}

	sort.Slice(y, func(i, j int) bool {
		return y[i] < y[j]
	})

	for i := 0; i < N; i++ {
		res := y[i] + float64(2*R*i)/100.0
		if i < N-1 {
			putf("%.2f ", res)
		} else {
			putf("%.2f\n", res)
		}
	}
}

func height(H0 float64, T int) float64 {
	if T < 0 {
		return H0
	}

	g := 10.0
	t := math.Sqrt(2.0 * H0 / g)

	T2 := float64(T) - math.Floor(float64(T)/t)*t
	if fall := int(math.Floor(float64(T)/t))%2 == 0; fall {
		return H0 - 0.5*g*T2*T2
	}
	return math.Sqrt(2.0*H0*g) - g*T2
}
