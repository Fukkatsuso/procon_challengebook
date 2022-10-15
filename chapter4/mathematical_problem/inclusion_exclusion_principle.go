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

func getInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = getInt()
	}
	return slice
}

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := getInt(), getInt()
	a := getInts(m)

	res := 0
	for i := 1; i < (1 << m); i++ {
		// iの立っているbitの個数
		// 包除原理の足し引き判断に使用
		num := 0
		for j := i; j != 0; j >>= 1 {
			num += j & 1
		}

		// iの立っているbitに対するa，すなわち {a[j] | (i>>j)&1==1} の最小公倍数
		lcm := 1
		for j := 0; j < m; j++ {
			if (i>>j)&1 == 1 {
				lcm = lcm / gcd(lcm, a[j]) * a[j]
				// lcmがnを超えると n/lcm = 0 なので，オーバーフローする前にbreak
				if lcm > n {
					break
				}
			}
		}

		if num%2 == 0 {
			res -= n / lcm
		} else {
			res += n / lcm
		}
	}
	puts(res)
}
