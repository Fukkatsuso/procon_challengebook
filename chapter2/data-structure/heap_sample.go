package main

import (
	"container/heap"
	"fmt"
)

type Elem struct {
	value    interface{}
	priority int
	index    int
}

type PriorityQueue []*Elem

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	elem := x.(*Elem)
	elem.index = n
	*pq = append(*pq, elem)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(*pq)
	elem := old[n-1]
	old[n-1] = nil
	elem.index = -1
	*pq = old[0 : n-1]
	return elem
}

func (pq *PriorityQueue) update(elem *Elem, value interface{}, priority int) {
	elem.value = value
	elem.priority = priority
	heap.Fix(pq, elem.index)
}

func main() {
	arr := []int{3, 2, 5, 1, 7, 9}
	pq := make(PriorityQueue, len(arr))
	for i := 0; i < len(arr); i++ {
		pq[i] = &Elem{
			value:    i,
			priority: arr[i],
			index:    i,
		}
	}
	heap.Init(&pq)

	elem := &Elem{
		value:    10,
		priority: 0,
	}
	heap.Push(&pq, elem)
	pq.update(elem, elem.value, 100)

	for pq.Len() > 0 {
		elem := heap.Pop(&pq).(*Elem)
		fmt.Println(elem.priority, elem.value)
	}
}
