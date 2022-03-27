package main

import (
	"fmt"

	. "github.com/ansurfen/go-with-stl/conf"
	. "github.com/ansurfen/go-with-stl/ds/queue"
	. "github.com/ansurfen/go-with-stl/utils"
)

type Mycircle[T Comparable] struct {
	cir *Circle[T]
}

func (m *Mycircle[T]) ForEach() {
	ForEach(m.cir.Iter().Data(), func(i T) {
		fmt.Print(i, " ")
	})
	fmt.Println()
}

func (m *Mycircle[T]) Size() int {
	return m.cir.Iter().Size()
}

func (m *Mycircle[T]) Enqueue(val ...T) {
	for _, v := range val {
		m.cir.Enqueue(v)
	}
}

func main() {
	var c Mycircle[int]
	c.cir = NewCircle([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	c.cir.SetSize(15)
	c.Enqueue(12, 13, 14, 15)
	c.cir.Dequeue()
	c.ForEach()
	fmt.Println(c.Size())
}
