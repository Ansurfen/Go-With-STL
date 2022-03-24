package main

import (
	"fmt"

	. "github.com/ansurfen/go-with-stl/algorithm"
	. "github.com/ansurfen/go-with-stl/utils"
)

func main() {
	it := NewIter([]int{9, 8, 7, 6, 5}, 10)
	BubbleSort(it.Data())
	for it.HasPrev() {
		fmt.Print(it.Value(), " ")
		it.Prev()
	}
	fmt.Println()
	ForEach(it.At(1).Data(), func(v int) {
		if v%2 == 0 {
			fmt.Print(v, " ")
		}
		if v == it.End().Value() {
			fmt.Println()
		}
	})
	it.Clear()
	fmt.Println(it)
}
