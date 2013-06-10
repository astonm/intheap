// Example integer heap using the built-in container/heap based on example
package main

import "container/heap"
import "fmt"

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

func main() {
	h := &IntHeap{}
	heap.Init(h)
	heap.Push(h, 2)
	heap.Push(h, 5)
	heap.Push(h, 3)
	heap.Push(h, 1)
	// heap.Push(h, "not a number") // this line will blow up at run time, but not compile time

	heap.Remove(h, 1)
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Println()
}
