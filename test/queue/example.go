package main

import (
	"fmt"

	. "github.com/ansurfen/go-with-stl/conf"
	. "github.com/ansurfen/go-with-stl/ds/queue"
	. "github.com/ansurfen/go-with-stl/utils"
)

type Myqueue[T Comparable] struct {
	que *Queue[T]
}

func (m *Myqueue[T]) Push(val ...T) {
	m.que.Push(val...)
}

func (m *Myqueue[T]) Pop(cnt ...int) {
	if len(cnt) == 0 {
		m.que.Pop()
	} else {
		for i := 0; i < cnt[0]; i++ {
			m.que.Pop()
		}
	}
}

func (m *Myqueue[T]) Eachfor() {
	ForEach(m.que.Iter().Data(), func(i T) {
		fmt.Print(i, " ")
	})
	fmt.Println()
}

func (m *Myqueue[T]) Size() int {
	return m.que.Iter().Size()
}

func (m *Myqueue[T]) Cap() int {
	return m.que.Iter().Cap()
}

func main() {
	var q Myqueue[int]
	q.que = NewQueue([]int{1, 2, 3})
	q.Push(4, 5, 6, 7)
	q.Pop(5)
	q.Eachfor()
	q.Push(8, 9, 10, 11, 12, 13)
	q.Pop(5)
	q.Eachfor()
	q.que.Release()
	fmt.Println(q.que)
}
