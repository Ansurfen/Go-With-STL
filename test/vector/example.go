package main

import (
	"fmt"
	. "github.com/ansurfen/go-with-stl/algorithm"
	. "github.com/ansurfen/go-with-stl/conf"
	. "github.com/ansurfen/go-with-stl/ds"
	. "github.com/ansurfen/go-with-stl/utils"
)

type Myvector[T Comparable] struct {
	vec *Vector[T]
}

func (m *Myvector[T]) Max() T {
	return Max(m.vec.Iter().Data())
}

func (m *Myvector[T]) Min() T {
	return Min(m.vec.Iter().Data())
}

func (m *Myvector[T]) ForEach() {
	ForEach(m.vec.Iter().Data(), func(i T) {
		fmt.Print(i, " ")
	})
}

func (m *Myvector[T]) Sort(sortFunc func(arr []T)) {
	sortFunc(m.vec.Iter().Data())
}

func (m *Myvector[T]) Reverse() {
	Reverse(m.vec.Iter().Data())
}

func main() {
	var m Myvector[int]
	m.vec = NewVector([]int{1, 6, 5})
	m.vec.Push_Back(1, 3, 5, 4, 7)
	m.Sort(InsertSort[int])
	m.vec.Pop_back()
	m.Reverse()
	m.ForEach()
}
