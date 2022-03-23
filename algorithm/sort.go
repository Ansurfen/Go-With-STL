package algorithm

import (
	. "github.com/ansurfen/go-with-stl/conf"
)

func BubbleSort[T Comparable](arr []T) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func InsertSort[T Comparable](arr []T) {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		pos := i
		for pos > 0 && arr[pos-1] > temp {
			arr[pos] = arr[pos-1]
			pos--
		}
		if pos != i {
			arr[pos] = temp
		}
	}
}

func SelectSort[T Comparable](arr []T) {
	var i, j int
	for i = 0; i < len(arr)-1; i++ {
		min := i
		for j = i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != j {
			arr[min], arr[i] = arr[i], arr[min]
		}
	}
}
