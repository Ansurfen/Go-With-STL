package main

import (
	"fmt"

	. "github.com/ansurfen/go-with-stl/conf"
	. "github.com/ansurfen/go-with-stl/ds/heap"
	. "github.com/ansurfen/go-with-stl/utils"
)

type Myheap[T Comparable] struct {
	heap *Heap[T]
}

func (m *Myheap[T]) Typeof() {
	if m.heap.Typeof() {
		fmt.Println("The heap is MAXHEAP")
	} else {
		fmt.Println("The heap is MINHEAP")
	}
}

func (m *Myheap[T]) ForEach() {
	ForEach(m.heap.Iter().Data(), func(i T) {
		fmt.Print(i, " ")
	})
	fmt.Println()
}

func (m *Myheap[T]) Size() int {
	return m.heap.Iter().Size()
}

func (m *Myheap[T]) Cap() int {
	return m.heap.Iter().Cap()
}

func main() {
	var h Myheap[int]
	h.heap = NewHeap([]int{1, 3, 2})
	h.heap.Push(6, 5, 7, 8, 9, 10, -1, -2)
	h.heap.SetType(MINHEAP)
	h.heap.Pop()
	h.Typeof()
	h.ForEach()
}
