package algorithm

import (
	. "github.com/ansurfen/go-with-stl/conf"
)

type Convertable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

//将要要支持区域
func LinerSearch[T Comparable](arr []T, val T) (index int) {
	index = -1
	if arr == nil || len(arr) == 0 {
		return index
	}
	for i := 0; i < len(arr); i++ {
		if val == arr[i] {
			index = i
			break
		}
	}
	return index
}

func BinarySearch[T Comparable](arr []T, val T) (index int) {
	index = -1
	if arr == nil || len(arr) == 0 {
		return index
	}
	begin, end := 0, len(arr)
	for begin < end {
		mid := (begin + end) >> 1
		if val < arr[mid] {
			end = mid
		} else if val > arr[mid] {
			begin = mid + 1
		} else {
			index = mid
			break
		}
	}
	return index
}

func InterpolationSearch[T Convertable](arr []T, val T) (index int) {
	index = -1
	if arr == nil || len(arr) == 0 {
		return index
	}
	begin, end := 0, len(arr)-1
	for begin < end {
		if arr[begin] == T(end) {
			return index
		}
		mid := begin + ((end - begin) * (int(val) - int(arr[begin])) / (int(arr[end]) - int(arr[begin])))
		if mid < begin || mid > end {
			return index
		}
		if val > arr[mid] {
			begin = mid + 1
		} else if val < arr[mid] {
			break
		} else {
			index = mid
			break
		}
	}
	return index
}
