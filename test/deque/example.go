package main

import (
	"fmt"

	. "github.com/ansurfen/go-with-stl/conf"
	. "github.com/ansurfen/go-with-stl/ds/queue"
	. "github.com/ansurfen/go-with-stl/utils"
)

type Mydeque[T Comparable] struct {
	deq *Deque[T]
}

func (m *Mydeque[T]) ForEach() {
	ForEach(m.deq.Iter().Data(), func(i T) {
		fmt.Print(i, " ")
	})
	fmt.Println()
}

func (m *Mydeque[T]) Size() int {
	return m.deq.Iter().Size()
}

func (m *Mydeque[T]) Push_front(val ...T) {
	for _, v := range val {
		m.deq.Push_front(v)
	}
}

func (m *Mydeque[T]) Push_back(val ...T) {
	for _, v := range val {
		m.deq.Push_back(v)
	}
}

func (m *Mydeque[T]) Pop_front(cnt ...int) {
	if len(cnt) == 0 {
		m.deq.Pop_front()
	} else {
		for i := 0; i < cnt[0]; i++ {
			m.deq.Pop_front()
		}
	}
}

func (m *Mydeque[T]) Pop_back(cnt ...int) {
	if len(cnt) == 0 {
		m.deq.Pop_back()
	} else {
		for i := 0; i < cnt[0]; i++ {
			m.deq.Pop_back()
		}
	}
}

func main() {
	var d Mydeque[int]
	d.deq = NewDeque([]int{})
	d.Push_front(3, 4, 5, 6, 7)
	d.Push_back(2, 1)
	d.Pop_back(2)
	d.Pop_front(2)
	d.ForEach()
	fmt.Println(d.Size())
}
