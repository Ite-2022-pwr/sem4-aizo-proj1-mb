package utils

import "golang.org/x/exp/constraints"

func CheckSortedList[T constraints.Ordered](arr []T) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
