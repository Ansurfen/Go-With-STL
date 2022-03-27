package heap

import (
	. "github.com/ansurfen/go-with-stl/conf"
	. "github.com/ansurfen/go-with-stl/utils"
)

const (
	MAXHEAP = true
	MINHEAP = false
)

type AbstractHeap[T Comparable] interface {
	Iter() *Iterator[T]
	Typeof() bool
	Fixed()
	Push()
	Pop()
	SetType(bool)
}

type Heap[T Comparable] struct {
	data     []T
	length   int
	heapType bool
}

func NewHeap[T Comparable](arr []T) *Heap[T] {
	if len(arr) == 0 {
		return &Heap[T]{
			data:     arr,
			length:   0,
			heapType: MAXHEAP,
		}
	}
	h := &Heap[T]{
		data:     arr,
		length:   len(arr),
		heapType: MAXHEAP,
	}
	h.Fixed()
	return h
}

func (h *Heap[T]) Iter() (i *Iterator[T]) {
	if h == nil {
		h = NewHeap([]T{})
	}
	i = NewIter(h.data)
	return i
}

func (h *Heap[T]) SetType(heapType bool) {
	h.heapType = heapType
	h.Fixed()
}

func (h *Heap[T]) Typeof() bool {
	return h.heapType
}

func (h *Heap[T]) Fixed() {
	for i := h.length / 2; i >= 0; i-- {
		h.sink(i)
	}
}

func (h *Heap[T]) sink(index int) {
	cmpVal := 1
	if !h.heapType {
		cmpVal = -1
	}
	if !h.hasChild(index) {
		return
	}
	temp := h.data[index]
	childIndex := 2*index + 1
	for h.hasChild(index) {
		if childIndex+1 < h.length && UpCompare(h.data[childIndex+1], h.data[childIndex]) == cmpVal {
			childIndex++
		}
		if temp == h.data[childIndex] || UpCompare(temp, h.data[childIndex]) == cmpVal {
			break
		}
		h.data[index] = h.data[childIndex]
		index = childIndex
		childIndex = 2*childIndex + 1
	}
	h.data[index] = temp
}

func (h *Heap[T]) float(index int) {
	cmpVal := 1
	if !h.heapType {
		cmpVal = -1
	}
	parentIndex := (index - 1) / 2
	temp := h.data[index]
	for index > 0 && UpCompare(temp, h.data[parentIndex]) == cmpVal {
		h.data[index] = h.data[parentIndex]
		index = parentIndex
		parentIndex = (parentIndex - 1) / 2
	}
	h.data[index] = temp
}

func (h *Heap[T]) Push(val ...T) {
	h.data = append(h.data, val...)
	for i := 0; i < len(val); i++ {
		h.float(h.length + i)
	}
	h.length = len(h.data)
}

func (h *Heap[T]) Pop() {
	h.data[0], h.data[h.length-1] = h.data[h.length-1], h.data[0]
	h.sink(0)
	h.data = h.data[:h.length-1]
	h.length--
}

func (h *Heap[T]) hasChild(index int) bool {
	if index*2+1 <= h.length-1 {
		return true
	}
	return false
}
