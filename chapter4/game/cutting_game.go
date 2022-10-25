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

	w, h := getInt(), getInt()

	// メモ用の配列
	mem := make([][]int, w+1)
	for i := range mem {
		mem[i] = make([]int, h+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	if grundy(w, h, mem) != 0 {
		puts("WIN")
	} else {
		puts("LOSE")
	}
}

func grundy(w, h int, mem [][]int) int {
	if mem[w][h] != -1 {
		return mem[w][h]
	}

	s := map[int]int{}
	// 横幅wの紙を縦に切って横幅i, w-iの紙にする
	for i := 2; w-i >= 2; i++ {
		s[grundy(i, h, mem)^grundy(w-i, h, mem)]++
	}
	// 縦幅hの紙を横に切って縦幅i, h-iの紙にする
	for i := 2; h-i >= 2; i++ {
		s[grundy(w, i, mem)^grundy(w, h-i, mem)]++
	}

	res := 0
	for s[res] > 0 {
		res++
	}
	mem[w][h] = res
	return res
}
