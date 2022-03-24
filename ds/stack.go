package ds

import (
	. "github.com/ansurfen/go-with-stl/conf"
	. "github.com/ansurfen/go-with-stl/utils"
)

type AbstractStack[T Comparable] interface {
	Iter() *Iterator[T]
	Size() int
	Clear()
	Empty() bool
	Push() T
	Pop()
	Top() T
}

type Stack[T Comparable] struct {
	data []T
	top  int
	cap  int
}

func NewStack[T Comparable](data []T) *Stack[T] {
	return &Stack[T]{
		data: data,
		top:  len(data) - 1,
		cap:  cap(data),
	}
}

func (s *Stack[T]) Iter() (i *Iterator[T]) {
	if s == nil {
		s = NewStack([]T{})
	}
	i = NewIter(s.data)
	return i
}

func (s *Stack[T]) Size() int {
	return s.top + 1
}

func (s *Stack[T]) Clear() {
	s.data = make([]T, 0)
	s.top = -1
	s.cap = 1
}

func (s *Stack[T]) Empty() bool {
	return s.top == -1
}

func (s *Stack[T]) Push(val T) {
	s.data = append(s.data, val)
	s.flashParam()
}

func (s *Stack[T]) Pop() {
	if s.Empty() {
		panic("")
	}
	s.data = s.data[:s.top]
	s.flashParam()
}

func (s *Stack[T]) Top() T {
	if s.Empty() {
		panic("")
	}
	return s.data[s.top]
}

func (s *Stack[T]) flashParam() {
	s.top = len(s.data) - 1
	s.cap = cap(s.data)
}
