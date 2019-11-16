package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

// 大きい順にヒープを作る
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	var n, l, p int
	fmt.Scan(&n, &l, &p)
	a := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	a[n] = l
	b := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}

	h := &IntHeap{}
	heap.Init(h)
	n++
	ans := 0
	pos := 0
	tank := p
	for i := 0; i < n; i++ {
		d := a[i] - pos
		for tank-d < 0 {
			if h.Len() == 0 {
				fmt.Println("-1")
				return
			}
			tank += heap.Pop(h).(int)
			ans++
		}
		tank -= d
		pos = a[i]
		heap.Push(h, b[i])
	}

	fmt.Println(ans)
}

// 4 25 10
// 10 14 20 21
// 10 5 2 4

//-> 2
