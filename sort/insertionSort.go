package sort

import "golang.org/x/exp/constraints"

func InsertionSort[T constraints.Ordered](list []T) []T {
	sortedList := make([]T, len(list))
	copy(sortedList, list)
	for i := 0; i < len(sortedList); i++ {
		key := sortedList[i]
		j := i - 1
		for j >= 0 && sortedList[j] > key {
			sortedList[j+1] = sortedList[j]
			j = j - 1
		}
		sortedList[j+1] = key
	}
	return sortedList
}
