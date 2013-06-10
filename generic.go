// Example integer heap using generic/heap
package main

import "github.com/astonm/generic/heap"
import "fmt"

type IntHeap struct {
	heap.GenericHeap

	Push   func(int)
	Pop    func() int
	Remove func(int) int
}

func (*IntHeap) Less(a, b int) bool { return a < b }

func main() {
	h := &IntHeap{}
	heap.Init(h)
	h.Push(2)
	h.Push(5)
	h.Push(3)
	h.Push(1)
	//h.Push("not a number") // this line will blow up at compile time!
	h.Remove(1)
	for h.Len() > 0 {
		fmt.Printf("%d ", h.Pop())
	}
	fmt.Println()
}
