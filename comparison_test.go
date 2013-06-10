package test

import "container/heap"
import generic_heap "github.com/astonm/generic/heap"
import "testing"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
        return x
}

type GenericIntHeap struct {
	generic_heap.GenericHeap

	Push func(int)
	Pop  func() int
	Remove func(int) int
}
func (*GenericIntHeap) Less(a, b int) bool { return a < b}

func newHeap() *IntHeap {
	h := &IntHeap{}
	heap.Init(h)
	return h
}

func newGenericHeap() *GenericIntHeap {
	h := &GenericIntHeap{}
	generic_heap.Init(h)
	return h
}

func BenchmarkHeap(b *testing.B) {
	h := newHeap()
	for i := 0; i < b.N; i++ {
		heap.Push(h, i)
	}
	for h.Len() > 0 {
		heap.Pop(h)
	}
}

func BenchmarkGenericHeap(b *testing.B) {
	h := newGenericHeap()
	for i := 0; i < b.N; i++ {
		h.Push(i)
	}
	for h.Len() > 0 {
		h.Pop()
	}
}
