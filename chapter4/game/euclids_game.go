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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	a, b := getInt(), getInt()

	f := true
	for {
		if a > b {
			a, b = b, a
		}

		// bがaの倍数なら勝ち
		if b%a == 0 {
			break
		}

		if b-a > a {
			break
		}

		b -= a
		f = !f
	}

	if f {
		puts("Stan wins")
	} else {
		puts("Ollie wins")
	}
}
