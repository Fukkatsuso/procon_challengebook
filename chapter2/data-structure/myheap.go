package main

import "fmt"

type Heap struct {
	size  int
	elems [1000]int
}

type MyHeap interface {
	Size() int
	Push(int)
	Pop() int
}

func (heap *Heap) Size() int {
	return heap.size
}

func (heap *Heap) Push(x int) {
	i := heap.size

	for i > 0 {
		// 親ノードの番号
		p := (i - 1) / 2

		// 逆転していないなら抜ける
		if heap.elems[p] <= x {
			break
		}

		// 親ノードの数字を下ろして、自分は上に
		heap.elems[i] = heap.elems[p]
		i = p
	}

	heap.elems[i] = x
	heap.size++
}

func (heap *Heap) Pop() int {
	// 最小値
	ret := heap.elems[0]
	heap.size--
	// rootに持ってくる値
	x := heap.elems[heap.size]

	// rootから下ろしていく
	i := 0
	for i*2+1 < heap.size {
		// 子同士を比較
		a, b := i*2+1, i*2+2
		if b < heap.size && heap.elems[b] < heap.elems[a] {
			a = b
		}
		if heap.elems[a] >= x {
			break
		}

		// 子の数字を持ち上げる
		heap.elems[i] = heap.elems[a]
		i = a
	}

	heap.elems[i] = x

	return ret
}

func main() {
	h := new(Heap)
	fmt.Println(h.Size())
	for i := 10; i < 20; i++ {
		h.Push(i)
	}
	fmt.Println(h.Pop())
}
