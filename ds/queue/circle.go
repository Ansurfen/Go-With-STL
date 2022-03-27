package queue

import (
	. "github.com/ansurfen/go-with-stl/conf"
	. "github.com/ansurfen/go-with-stl/utils"
)

type AbstractCircle[T Comparable] interface {
	Iter() *Iterator[T]
	Enqueue()
	Dequeue()
	SetSize(int)
	Front() T
	Back() T
}

type Circle[T Comparable] struct {
	head *cirNode[T]
	tail *cirNode[T]
	cap  int
	size int
}

type cirNode[T Comparable] struct {
	data T
	next *cirNode[T]
}

func NewCircle[T Comparable](data []T) *Circle[T] {
	c := &Circle[T]{
		head: nil,
		tail: nil,
		cap:  10,
		size: 0,
	}
	if len(data) == 0 {
		return c
	}
	for _, v := range data {
		c.Enqueue(v)
	}
	return c
}

func newCirNode[T Comparable](data T) *cirNode[T] {
	return &cirNode[T]{
		data: data,
		next: nil,
	}
}

func (c *Circle[T]) Iter() (i *Iterator[T]) {
	if c == nil {
		c = NewCircle([]T{})
	}
	tmp := make([]T, c.size)
	cnt := 0
	for n := c.head; cnt < c.size; n = n.next {
		tmp[cnt] = n.data
		cnt++
	}
	i = NewIter(tmp)
	return i
}

func (c *Circle[T]) Enqueue(val T) {
	n := newCirNode(val)
	if c.size == 0 {
		c.head = n
		c.tail = n
		n.next = n
	} else {
		if c.size < c.cap {
			t := c.head
			for t != c.tail {
				t = t.next
			}
			n.next = t.next
			t.next = n
			c.tail = n
		} else {
			c.tail = c.tail.next
			c.tail.data = val
			return
		}
	}
	c.size++
}

func (c *Circle[T]) Dequeue() {
	if c.size == 0 {
		c = NewCircle([]T{})
	}
	c.head = c.head.next
	c.size--
}

func (c *Circle[T]) SetSize(size int) {
	if size <= 0 {
		size = 1
	}
	if size < c.cap {
		if size < c.size {
			n := c.head
			cnt := 0
			for ; cnt < size; n = n.next {
				cnt++
			}
			n.next = c.tail
		}
		c.size = size
	}
	c.cap = size
}

func (c *Circle[T]) Front() T {
	return c.head.data
}

func (c *Circle[T]) Back() T {
	return c.tail.data
}
