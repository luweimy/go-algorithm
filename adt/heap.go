package main

import (
	"log"
	"math/rand"
	"time"
)

// 大根堆
// https://www.cnblogs.com/wangchaowei/p/8288216.html
// https://www.jianshu.com/p/21bef3fc3030
type MaxHeap struct {
	heap []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (h *MaxHeap) Add(e int) {
	h.heap = append(h.heap, e)
	var cur = len(h.heap) - 1
	for cur > 0 {
		parent := h.parent(cur)
		if h.less(cur, parent) {
			break
		}
		h.swap(cur, parent)
		cur = parent
	}
}

func (h *MaxHeap) Pop() int {
	if len(h.heap) <= 0 {
		panic(-1)
	}
	var r = h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]

	for cur := 0; cur < len(h.heap); {
		l, r := h.left(cur), h.right(cur)
		if l > len(h.heap)-1 {
			break
		}
		var max = l
		if r < len(h.heap)-1 && h.less(l, r) {
			max = r
		}
		if h.less(max, cur) {
			break
		}
		h.swap(max, cur)
		cur = max
	}
	return r
}

func (h *MaxHeap) Len() int {
	return len(h.heap)
}

func (h *MaxHeap) less(i, j int) bool {
	return h.heap[i] < h.heap[j]
}

func (h *MaxHeap) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *MaxHeap) left(i int) int {
	return 2*i + 1
}

func (h *MaxHeap) right(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) parent(i int) int {
	return (i - 1) / 2
}

func main() {
	var ints = randomInts(100)
	var mh = NewMaxHeap()
	for _, e := range ints {
		mh.Add(e)
	}
	log.Println("heap=>", mh.heap)

	// 测试结果
	var rt = make([]int, 0, len(ints))
	for mh.Len() > 0 {
		rt = append(rt, mh.Pop())
	}
	log.Println("result=>", intsAreDescSorted(rt), rt)
}

func randomInts(n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = rand.Intn(n)
	}
	return res
}

func intsAreDescSorted(s []int) bool {
	for i := 1; i < len(s); i++ {
		if s[i] > s[i-1] {
			return false
		}
	}
	return true
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
