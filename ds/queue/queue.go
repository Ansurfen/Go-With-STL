package queue

import (
	. "github.com/ansurfen/go-with-stl/conf"
	. "github.com/ansurfen/go-with-stl/utils"
)

type AbstractQueue[T Comparable] interface {
	Iter() *Iterator[T]
	Clear()
	Empty() bool
	Front() T
	Back() T
	Push()
	Pop() T
}

type Queue[T Comparable] struct {
	data         []T
	head         int
	rear         int
	offset       int
	shrinkFactor int
}

func NewQueue[T Comparable](data []T) *Queue[T] {
	return &Queue[T]{
		data:         data,
		head:         0,
		rear:         len(data) - 1,
		offset:       0,
		shrinkFactor: 5,
	}
}

func (q *Queue[T]) Iter() (i *Iterator[T]) {
	if q == nil {
		q = NewQueue([]T{})
	}
	i = NewIter(q.data[q.head-q.offset : q.rear-q.offset+1])
	return i
}

func (q *Queue[T]) Clear() {
	q.data = q.data[:0]
	q.offset = 0
	q.flashParam()
}

func (q *Queue[T]) Release() {
	q.data = q.data[q.head-q.offset:]
	q.offset = 0
	q.head = 0
	q.flashParam()
}

func (q *Queue[T]) SetFactor(val int) {
	if val <= 0 {
		val = 1
	}
	q.shrinkFactor = val
}

func (q *Queue[T]) Empty() bool {
	return q.rear == -1
}

func (q *Queue[T]) Front() T {
	if !q.Empty() {
		return q.data[q.head-q.offset]
	} else {
		panic("The queue is empty")
	}
}

func (q *Queue[T]) Back() T {
	if !q.Empty() {
		return q.data[q.rear-q.offset]
	} else {
		panic("The queue is empty")
	}
}

func (q *Queue[T]) Push(val ...T) {
	q.data = append(q.data, val...)
	for i := 0; i < len(val); i++ {
		q.rear++
	}
}

func (q *Queue[T]) Pop() (val T) {
	val = q.data[q.head-q.offset]
	if q.head < q.rear {
		q.head++
	}
	if q.head%q.shrinkFactor == 0 {
		q.data = q.data[q.head-q.offset:]
		q.offset += q.shrinkFactor
	}
	return val
}

func (q *Queue[T]) flashParam() {
	q.rear = len(q.data) - 1
}
