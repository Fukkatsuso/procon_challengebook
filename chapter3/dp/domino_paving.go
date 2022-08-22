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
	M              = 1000000007 // 書籍中のコードにMの値がなかったので，よく使われる値を設定
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
	black := make([][]bool, n)
	for i := 0; i < n; i++ {
		line := gets()
		black[i] = make([]bool, m)
		for j := 0; j < m; j++ {
			black[i][j] = line[j] == 'x'
		}
	}

	current := make([]int, 1<<m)
	current[0] = 1

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			next := make([]int, 1<<m)
			for used := 0; used < (1 << m); used++ {
				if (used>>j)&1 == 1 || black[i][j] {
					next[used] = current[used&^(1<<j)]
				} else {
					res := 0
					// マス(i,j)と(i,j+1)に置く
					if j+1 < m && ((used>>(j+1))&1 == 0) && !black[i][j+1] {
						res += current[used|(1<<(j+1))]
					}
					// マス(i,j)と(i+1,j)に置く
					if i+1 < n && !black[i+1][j] {
						res += current[used|(1<<j)]
					}
					next[used] = res % M
				}
			}
			current = next
		}
	}

	puts(current[0])
}
