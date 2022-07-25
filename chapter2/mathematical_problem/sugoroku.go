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

func getInt2() (int, int) {
	return getInt(), getInt()
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

// ax + by = 1 の整数解(x, y)を求める
// 返り値のdは gcd(a, b)
func extGcd(a, b int) (d, x, y int) {
	d = a
	if b != 0 {
		d, y, x = extGcd(b, a%b)
		y -= a / b * x
	} else {
		x, y = 1, 0
	}
	return d, x, y
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	a, b := getInt2()

	if gcd(a, b) != 1 {
		puts(-1)
		return
	}

	_, x, y := extGcd(a, b)
	ans := [4]int{0, 0, 0, 0}
	if x > 0 {
		ans[0] = x
	} else {
		ans[2] = -x
	}
	if y > 0 {
		ans[1] = y
	} else {
		ans[3] = -y
	}
	puts(ans[0], ans[1], ans[2], ans[3])
}
