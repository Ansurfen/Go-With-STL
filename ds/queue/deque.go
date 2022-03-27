package queue

import (
	. "github.com/ansurfen/go-with-stl/conf"
	. "github.com/ansurfen/go-with-stl/utils"
)

type AbstractDeque[T Comparable] interface {
	Iter() *Iterator[T]
	Empty() bool
	Front() T
	Back() T
	Push_front(T)
	Push_back(T)
	Pop_front()
	Pop_back()
}

type Deque[T Comparable] struct {
	head *dequeNode[T]
	tail *dequeNode[T]
	size int
}

type dequeNode[T Comparable] struct {
	data T
	prev *dequeNode[T]
	next *dequeNode[T]
}

func NewDeque[T Comparable](arr []T) *Deque[T] {
	d := &Deque[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
	for _, v := range arr {
		d.Push_front(v)
	}
	return d
}

func newDequeNode[T Comparable](val T) *dequeNode[T] {
	return &dequeNode[T]{
		data: val,
		prev: nil,
		next: nil,
	}
}

func (d *Deque[T]) Iter() (i *Iterator[T]) {
	if d == nil {
		d = NewDeque([]T{})
	}
	cnt := 0
	tmp := make([]T, d.size)
	for n := d.head; cnt < d.size && n != nil; n = n.next {
		tmp[cnt] = n.data
		cnt++
	}
	i = NewIter(tmp)
	return i
}

func (d *Deque[T]) Push_front(val T) {
	if d == nil {
		d = NewDeque([]T{})
	}
	n := newDequeNode(val)
	if d.size == 0 {
		d.head = n
		d.tail = n
	} else {
		n.insertNext(d.head)
		d.head = n
	}
	d.size++
}

func (d *Deque[T]) Push_back(val T) {
	if d == nil {
		d = NewDeque([]T{})
	}
	n := newDequeNode(val)
	if d.size == 0 {
		d.head = n
		d.tail = n
	} else {
		n.insertPrev(d.tail)
		d.tail = n
	}
	d.size++
}

func (d *Deque[T]) Pop_back() {
	if d.size == 0 {
		d = NewDeque([]T{})
		return
	}
	d.tail = d.tail.prev
	d.size--
}

func (d *Deque[T]) Pop_front() {
	if d.size == 0 {
		d = NewDeque([]T{})
		return
	}
	d.head = d.head.next
	d.size--
}

func (n *dequeNode[T]) insertNext(next *dequeNode[T]) {
	if n == nil || next == nil {
		return
	}
	next.prev = n
	if n.next != nil {
		n.next.prev = next
	}
	n.next = next
}

func (n *dequeNode[T]) insertPrev(prev *dequeNode[T]) {
	if n == nil || prev == nil {
		return
	}
	prev.next = n
	if n.prev != nil {
		n.prev.next = prev
	}
	n.prev = prev
}

func (d *Deque[T]) Empty() bool {
	return d.size == 0
}

func (d *Deque[T]) Front() T {
	return d.head.data
}

func (d *Deque[T]) Back() T {
	return d.tail.data
}
