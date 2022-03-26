package main

import (
	"fmt"

	. "github.com/ansurfen/go-with-stl/conf"
	. "github.com/ansurfen/go-with-stl/ds/stack"
	. "github.com/ansurfen/go-with-stl/utils"
)

type Mystack[T Comparable] struct {
	stk *Stack[T]
}

func (m *Mystack[T]) Reverse() {
	Reverse(m.stk.Iter().Data())
}

func (m *Mystack[T]) Eachfor() {
	t := Clone(m)
	t.Reverse()
	ForEach(m.stk.Iter().Data(), func(i T) {
		fmt.Print(i, " ")
	})
	fmt.Println()
}

func (m *Mystack[T]) Max() T {
	return Max(m.stk.Iter().Data())
}

func (m *Mystack[T]) Size() int {
	return m.stk.Iter().Size()
}

func (m *Mystack[T]) Cap() int {
	return m.stk.Iter().Cap()
}

func main() {
	var m Mystack[float32]
	m.stk = NewStack([]float32{1.2, 8.5, 0.2, 1.7})
	m.stk.Push(1.1)
	fmt.Println(m.Max())
	m.stk.Pop()
	m.Eachfor()
	fmt.Println(m.Size())
}
