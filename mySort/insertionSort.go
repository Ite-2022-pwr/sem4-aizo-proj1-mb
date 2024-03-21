package mySort

import "golang.org/x/exp/constraints"

func InsertionSort[T constraints.Ordered](list []T) []T {
	for i := 0; i < len(list); i++ {
		key := list[i]
		j := i - 1
		for j >= 0 && list[j] > key {
			list[j+1] = list[j]
			j = j - 1
		}
		list[j+1] = key
	}
	return list
}
