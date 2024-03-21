package sort

import "golang.org/x/exp/constraints"

func heapify[T constraints.Ordered](list []T, n, i int) {
	var maximum = i
	var l = 2*i + 1
	var r = 2*i + 2
	if l < n && list[maximum] < list[l] {
		maximum = l
	}
	if r < n && list[maximum] < list[r] {
		maximum = r
	}
	if maximum != i {
		list[maximum], list[i] = list[i], list[maximum]
		heapify(list, n, maximum)
	}
}

func HeapSort[T constraints.Ordered](list []T) []T {
	n := len(list)
	sortedList := make([]T, n)
	copy(sortedList, list)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(sortedList, n, i)
	}
	for i := n - 1; i > 0; i-- {
		sortedList[0], sortedList[i] = sortedList[i], sortedList[0]
		heapify(sortedList, i, 0)
	}
	return sortedList
}
