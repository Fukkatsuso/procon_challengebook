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
	inf            = 1 << 60
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

func getInt3() (int, int, int) {
	return getInt(), getInt(), getInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, ml, md := getInt3()
	al, bl, dl := make([]int, ml), make([]int, ml), make([]int, ml)
	for i := 0; i < ml; i++ {
		al[i], bl[i], dl[i] = getInt()-1, getInt()-1, getInt()
	}
	ad, bd, dd := make([]int, md), make([]int, md), make([]int, md)
	for i := 0; i < md; i++ {
		ad[i], bd[i], dd[i] = getInt()-1, getInt()-1, getInt()
	}

	d := make([]int, n)

	// 負閉路チェック
	if Bellmanford(n, al, bl, dl, ad, bd, dd, d) {
		puts(-1)
		return
	}

	for i := range d {
		d[i] = inf
	}
	d[0] = 0
	Bellmanford(n, al, bl, dl, ad, bd, dd, d)
	ans := d[n-1]
	if ans == inf {
		ans = -2
	}
	puts(ans)
}

// ベルマンフォード法
func Bellmanford(n int, al, bl, dl, ad, bd, dd, d []int) bool {
	// 最終ループでも更新されていれば，閉路があると判定
	updated := false

	for k := 0; k <= n; k++ {
		updated = false

		// i+1からiへコスト0の辺
		for i := 0; i+1 < n; i++ {
			if d[i+1] < inf && d[i] > d[i+1] {
				d[i] = d[i+1]
				updated = true
			}
		}

		// alからblへコストdlの辺
		for i := 0; i < len(al); i++ {
			if d[al[i]] < inf && d[bl[i]] > d[al[i]]+dl[i] {
				d[bl[i]] = d[al[i]] + dl[i]
				updated = true
			}
		}

		// bdからadへコスト-ddの辺
		for i := 0; i < len(bd); i++ {
			if d[bd[i]] < inf && d[ad[i]] > d[bd[i]]-dd[i] {
				d[ad[i]] = d[bd[i]] - dd[i]
				updated = true
			}
		}
	}

	return updated
}
