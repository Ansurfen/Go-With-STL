package utils

import (
	. "github.com/ansurfen/go-with-stl/conf"
)

type AbstractIter[T any] interface {
	Cap() int
	Size() int
	Value() T
	Data() []T
	Clear()
	Next() bool
	Prev() bool
	HasNext() bool
	HasPrev() bool
	At(int) *AbstractIter[T]
	Begin() *AbstractIter[T]
	End() *AbstractIter[T]
}

type Iterator[T any] struct {
	data  []T
	index int
}

func NewIter[T any](arr []T, index ...int) *Iterator[T] {
	var i int
	if index == nil {
		i = -1
	} else {
		i = index[0]
	}
	return &Iterator[T]{
		data:  arr,
		index: i,
	}
}

func (i *Iterator[T]) Begin() *Iterator[T] {
	if i == nil {
		return nil
	}
	return &Iterator[T]{
		data:  i.data,
		index: i.index,
	}
}

func (i *Iterator[T]) End() *Iterator[T] {
	if i == nil {
		return nil
	}
	return &Iterator[T]{
		data:  i.data,
		index: len(i.data) - 1,
	}
}

func (i *Iterator[T]) At(index int) *Iterator[T] {
	if i == nil {
		return nil
	}
	if index < 0 {
		index = 0
	}
	if index >= len(i.data) {
		index = len(i.data) - 1
	}
	return &Iterator[T]{
		data:  i.data,
		index: index,
	}
}

func (i *Iterator[T]) Value() T {
	if i.index <= 0 {
		i.index = 0
	}
	if i.index >= len(i.data) {
		i.index = len(i.data) - 1
	}
	return i.data[i.index]
}

func (i *Iterator[T]) Clear() {
	i.data = i.data[:0]
	i.index = -1
}

func (i *Iterator[T]) HasNext() bool {
	if i == nil || len(i.data) == 0 || i.index+1 > len(i.data) {
		return false
	}
	return true
}

func (i *Iterator[T]) HasPrev() bool {
	if i == nil || len(i.data) == 0 || i.index-1 < -1 {
		return false
	}
	return true
}

func (i *Iterator[T]) Next() bool {
	if i.HasNext() {
		i.index++
		return true
	}
	return false
}

func (i *Iterator[T]) Prev() bool {
	if i.HasPrev() {
		i.index--
		return true
	}
	return false
}

func (i *Iterator[T]) Size() int {
	return len(i.data)
}

func (i *Iterator[T]) Cap() int {
	return cap(i.data)
}

func (i *Iterator[T]) Data() []T {
	if i.index <= 0 {
		i.index = 0
	}
	if i.index >= len(i.data) {
		i.index = len(i.data) - 1
	}
	return i.data[i.index:]
}

func Clone[T any](src T) T {
	return src
}

func ForEach[T any](arr []T, handler func(T)) {
	for _, v := range arr {
		handler(v)
	}
}

func Reverse[T Comparable](arr []T) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}
