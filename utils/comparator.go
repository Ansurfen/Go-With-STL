package utils

import . "github.com/ansurfen/go-with-stl/conf"

func Min[T Comparable](arr []T) T {
	min := arr[0]
	for _, v := range arr {
		if min > v {
			min = v
		}
	}
	return min
}

func Max[T Comparable](arr []T) T {
	max := arr[0]
	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	return max
}

func Compare[T Comparable](a, b T) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}
