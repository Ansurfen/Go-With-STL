package ds

import (
	. "github.com/ansurfen/go-with-stl/conf"
	. "github.com/ansurfen/go-with-stl/utils"
)

type AbstractVec[T Comparable] interface {
	Size() int
	Cap() int
	Empty() bool
	Clear()
	Push_Back()
	Pop_back()
	Begin() *Iterator[T]
	End() *Iterator[T]
	InsertAt()
	RemoveAt()
	Erase()
	Reverse()
	Resize()
}

type Vector[T Comparable] struct {
	data    []T
	length  int
	capcity int
}

func (v *Vector[T]) Iter() (i *Iterator[T]) {
	if v == nil {
		v = NewVector([]T{})
	}
	i = NewIter(v.data)
	return i
}

func NewVector[T Comparable](data []T) *Vector[T] {
	if len(data) == 0 {
		return &Vector[T]{
			data:    data,
			length:  0,
			capcity: 1,
		}
	}
	return &Vector[T]{
		data:    data,
		length:  len(data),
		capcity: cap(data),
	}
}

func (v *Vector[T]) Push_Back(val ...T) {
	v.data = append(v.data, val...)
	v.flashParam()
}

func (v *Vector[T]) Pop_back() {
	v.data = v.data[:v.length-1]
	v.flashParam()
}

func (v *Vector[T]) Begin() *Iterator[T] {
	return v.Iter().Begin()
}

func (v *Vector[T]) End() *Iterator[T] {
	return v.Iter().End()
}

func (v *Vector[T]) At(index int) *Iterator[T] {
	return v.Iter().At(index)
}

func (v *Vector[T]) Size() int {
	return v.length
}

func (v *Vector[T]) Cap() int {
	return v.capcity
}

func (v *Vector[T]) Empty() bool {
	return v.length == 0
}

func (v *Vector[T]) Clear() {
	v.data = v.data[:0]
	v.flashParam()
}

func (v *Vector[T]) InsertAt(index int, t T) {
	v.data = append(v.data, t)
	v.flashParam()
	v.data[index], v.data[v.length-1] = v.data[v.length-1], v.data[index]
}

func (v *Vector[T]) RemoveAt(index int) {
	v.data = append(v.data[:index], v.data[index+1:]...)
	v.flashParam()
}

func (v *Vector[T]) Erase(first int, last int) {
	v.data = append(v.data[:first], v.data[last:]...)
	v.flashParam()
}

func (v *Vector[T]) Resize(size int) {
	if v.length > size {
		v.data = v.data[:size]
		v.flashParam()
	}
}

func (v *Vector[T]) flashParam() {
	v.length = len(v.data)
	v.capcity = cap(v.data)
}
