package main

import (
	"fmt"

	. "github.com/ansurfen/go-with-stl/conf"
	. "github.com/ansurfen/go-with-stl/utils"
)

type Mycmp[T Comparable] struct {
	it  *Iterator[T]
	cmp Comparator[T, int]
}

func NewMycmp[T Comparable, ret any](data []T, cmps ...Comparator[T, int]) *Mycmp[T] {
	var cmp Comparator[T, int]
	if len(cmps) == 0 {
		cmp = UpCompare[T]
	} else {
		cmp = cmps[0]
	}
	return &Mycmp[T]{
		it:  NewIter(data),
		cmp: cmp,
	}
}

func (m *Mycmp[T]) ForEach() {
	ForEach(m.it.Data(), func(i T) {
		fmt.Print(i, " ")
	})
	fmt.Println()
}

func (m *Mycmp[T]) Sort() {
	arr := m.it.Data()
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if m.cmp(arr[j], arr[j+1]) == 1 {
				tmp := arr[j]
				m.it.SetValue(j, arr[j+1])
				m.it.SetValue(j+1, tmp)
			}
		}
	}
}

func main() {
	c := NewMycmp[int, int]([]int{1, 3, 2, 4})
	c.Sort()
	c.ForEach()
}
