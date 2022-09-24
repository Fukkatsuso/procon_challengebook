package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func getFloat64() float64 {
	f, _ := strconv.ParseFloat(gets(), 64)
	return f
}

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	N := getInt()
	x, y, r := make([]float64, N), make([]float64, N), make([]float64, N)
	for i := 0; i < N; i++ {
		x[i], y[i], r[i] = getFloat64(), getFloat64(), getFloat64()
	}

	// 円iが円jの内側にあるか
	inside := func(i, j int) bool {
		dx, dy := x[i]-x[j], y[i]-y[j]
		return dx*dx+dy*dy <= r[j]*r[j]
	}

	// 円の左端と右端のx座標
	type Event struct {
		point float64
		id    int
	}
	events := make([]Event, 0)
	for i := 0; i < N; i++ {
		// 左端
		events = append(events, Event{
			point: x[i] - r[i],
			id:    i,
		})
		// 右端
		events = append(events, Event{
			point: x[i] + r[i],
			id:    i + N,
		})
	}
	sort.Slice(events, func(i, j int) bool {
		if events[i].point == events[j].point {
			return events[i].id < events[j].id
		}
		return events[i].point < events[j].point
	})

	// 平面走査
	outers := NewTree() // 走査線と交わっていて最も左側にある円の一覧
	res := make([]int, 0)
	for _, event := range events {
		id := event.id % N
		if event.id < N { // 円の左端
			// 円idの上に一番近く位置する円と比較
			lb := outers.LowerBound(Outer{
				y:  y[id],
				id: id,
			})
			if lb != nil && inside(id, lb.item.(Outer).id) {
				continue
			}
			// 円idの下に一番近く位置する円と比較
			mu := outers.MaxUnder(Outer{
				y:  y[id],
				id: id,
			})
			if mu != nil && inside(id, mu.item.(Outer).id) {
				continue
			}

			res = append(res, id+1)
			outers.Insert(Outer{
				y:  y[id],
				id: id,
			})
		} else { // 円の右端
			outers.Delete(Outer{
				y:  y[id],
				id: id,
			})
		}
	}

	sort.Ints(res)
	puts(len(res))
	for i := range res {
		if i == len(res)-1 {
			putf("%d\n", res[i])
		} else {
			putf("%d ", res[i])
		}
	}
}

type Outer struct {
	y  float64
	id int
}

func (outer Outer) Eq(x Item) bool {
	return outer.y == x.(Outer).y && outer.id == x.(Outer).id
}

func (outer Outer) Less(x Item) bool {
	if outer.y == x.(Outer).y {
		return outer.id < x.(Outer).id
	}
	return outer.y < x.(Outer).y
}

// 二分木
// 参考：
// http://www.nct9.ne.jp/m_hiroi/golang/abcgo10.html
// https://qiita.com/pandachan5228/items/127123037abf341e324d

// 二分木
type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return new(Tree)
}

func (t *Tree) Search(x Item) *Node {
	return searchNode(t.root, x)
}

func (t *Tree) Insert(x Item) {
	t.root = insertNode(t.root, x)
}

func (t *Tree) Delete(x Item) {
	t.root = deleteNode(t.root, x)
}

func (t *Tree) LowerBound(x Item) *Node {
	return lowerBoundNode(t.root, x)
}

func lowerBoundNode(n *Node, x Item) *Node {
	if n == nil {
		return nil
	}

	if n.item.Less(x) {
		if n.right != nil {
			n = lowerBoundNode(n.right, x)
		} else {
			return nil
		}
	} else {
		if n.left != nil && !n.left.item.Less(x) {
			n = lowerBoundNode(n.left, x)
		}
	}
	return n
}

func (t *Tree) UpperBound(x Item) *Node {
	return upperBoundNode(t.root, x)
}

func upperBoundNode(n *Node, x Item) *Node {
	if n == nil {
		return nil
	}

	if n.item.Less(x) || n.item.Eq(x) {
		if n.right != nil {
			n = upperBoundNode(n.right, x)
		} else {
			return nil
		}
	} else {
		if n.left != nil && !n.left.item.Less(x) && !n.left.item.Eq(x) {
			n = upperBoundNode(n.left, x)
		}
	}
	return n
}

// xより小さいもののうち最大のItem
func (t *Tree) MaxUnder(x Item) *Node {
	return maxUnder(t.root, x)
}
func maxUnder(n *Node, x Item) *Node {
	if n == nil {
		return nil
	}

	if !n.item.Less(x) {
		if n.left != nil {
			n = maxUnder(n.left, x)
		} else {
			return nil
		}
	} else {
		if n.right != nil && n.right.item.Less(x) {
			n = maxUnder(n.right, x)
		}
	}
	return n
}

// 格納するデータ
type Item interface {
	Eq(Item) bool
	Less(Item) bool
}

type Node struct {
	item        Item
	left, right *Node
}

func newNode(x Item) *Node {
	p := new(Node)
	p.item = x
	return p
}

func searchNode(node *Node, x Item) *Node {
	for node != nil {
		switch {
		case x.Eq(node.item):
			return node
		case x.Less(node.item):
			node = node.left
		default:
			node = node.right
		}
	}
	return nil
}

func insertNode(node *Node, x Item) *Node {
	switch {
	case node == nil:
		return newNode(x)
	case x.Eq(node.item):
		return node
	case x.Less(node.item):
		node.left = insertNode(node.left, x)
	default:
		node.right = insertNode(node.right, x)
	}
	return node
}

func searchMinNode(node *Node) Item {
	if node.left == nil {
		return node.item
	}
	return searchMinNode(node.left)
}

func deleteMinNode(node *Node) *Node {
	if node.left == nil {
		return node.right
	}
	node.left = deleteMinNode(node.left)
	return node
}

func searchMaxNode(node *Node) Item {
	if node.right == nil {
		return node.item
	}
	return searchMinNode(node.right)
}

func deleteNode(node *Node, x Item) *Node {
	if node != nil {
		if x.Eq(node.item) {
			if node.left == nil {
				return node.right
			} else if node.right == nil {
				return node.left
			} else {
				node.item = searchMinNode(node.right)
				node.right = deleteMinNode(node.right)
			}
		} else if x.Less(node.item) {
			node.left = deleteNode(node.left, x)
		} else {
			node.right = deleteNode(node.right, x)
		}
	}
	return node
}
