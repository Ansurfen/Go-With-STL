package main

import (
	"fmt"

	. "github.com/ansurfen/go-with-stl/conf"
	. "github.com/ansurfen/go-with-stl/ds"
	. "github.com/ansurfen/go-with-stl/utils"
)

type Myqueue[T Comparable] struct {
	que *Queue[T]
}

func (m *Myqueue[T]) Eachfor() {
	ForEach(m.que.Iter().Data(), func(i T) {
		fmt.Print(i, " ")
	})
	fmt.Println()
}

func main() {
	var q Myqueue[int]
	q.que = NewQueue([]int{1, 2, 3})
	q.que.Push(4, 5)
	q.Eachfor()
	fmt.Println(q.que.Pop())
	q.que.Pop()
	q.Eachfor()
	fmt.Println(q.que)
}
