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

type Point struct {
	x, y int
}

func isCrossing(p1, q1, p2, q2 Point) bool {
	s1 := (p1.x-q1.x)*(p2.y-p1.y) - (p1.y-q1.y)*(p2.x-p1.x)
	t1 := (p1.x-q1.x)*(q2.y-p1.y) - (p1.y-q1.y)*(q2.x-p1.x)

	s2 := (p2.x-q2.x)*(p1.y-p2.y) - (p2.y-q2.y)*(p1.x-p2.x)
	t2 := (p2.x-q2.x)*(q1.y-p2.y) - (p2.y-q2.y)*(q1.x-p2.x)

	return s1*t1 <= 0 && s2*t2 <= 0
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()
	p, q := make([]Point, n), make([]Point, n)
	for i := 0; i < n; i++ {
		x, y := getInt(), getInt()
		p[i] = Point{
			x: x,
			y: y,
		}
	}
	for i := 0; i < n; i++ {
		x, y := getInt(), getInt()
		q[i] = Point{
			x: x,
			y: y,
		}
	}

	m := getInt()
	for i := 0; i < m; i++ {
		a, b := getInt()-1, getInt()-1
		if isCrossing(p[a], q[a], p[b], q[b]) {
			puts("CONNECTED")
		} else {
			puts("NOT CONNECTED")
		}
	}
}
