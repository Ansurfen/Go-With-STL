package ds

import (
	. "github.com/ansurfen/go-with-stl/conf"
	. "github.com/ansurfen/go-with-stl/utils"
)

type AbstractQueue[T Comparable] interface {
	Iter() *Iterator[T]
	Size() int
	Clear()
	Empty() bool
	Front() T
	Back() T
	Push()
	Pop() T
}

type Queue[T Comparable] struct {
	data []T
	head int
	rear int
}

func NewQueue[T Comparable](data []T) *Queue[T] {
	return &Queue[T]{
		data: data,
		head: 0,
		rear: len(data),
	}
}

func (q *Queue[T]) Iter() (i *Iterator[T]) {
	if q == nil {
		q = NewQueue([]T{})
	}
	i = NewIter(q.data)
	return i
}

func (q *Queue[T]) Size() int {
	return q.rear - q.head
}

func (q *Queue[T]) Empty() bool {
	return q.Size() <= 0
}

func (q *Queue[T]) Clear() {
	q.data = q.data[:0]
	q.flashParam()
}

func (q *Queue[T]) flashParam() {
	q.rear = len(q.data)
}

func (q *Queue[T]) Push(val ...T) {
	q.data = append(q.data, val...)
	for i := 0; i < len(val); i++ {
		q.rear++
	}
}

func (q *Queue[T]) Pop() (ret T) {
	ret = q.data[q.head]
	q.data = append(q.data[:q.head], q.data[q.head+1:]...)
	q.rear--
	q.head = 0
	return ret
}

func (q *Queue[T]) Front() T {
	return q.data[q.head]
}

func (q *Queue[T]) Back() T {
	return q.data[q.rear]
}
