package sort

import (
	"golang.org/x/exp/constraints"
	"math"
)

func calculateGap(n, k, gapType int) int {
	gap := n
	switch gapType {
	case 0:
		gap = 2*int(math.Floor(float64(n)/math.Pow(2, float64(k+1)))) + 1
	case 1:
		gap = int(math.Floor(float64(n) / math.Pow(2, float64(k))))
	}

	return gap
}

func ShellSort[T constraints.Ordered](list []T, n, gapType int) []T {
	sortedList := make([]T, len(list))
	copy(sortedList, list)
	k, gap := 1, n
	for gap >= 1 {
		gap = calculateGap(n, k, gapType)
		k++
		for i := gap; i < n; i++ {
			temp := sortedList[i]
			j := i
			for j = j; j >= gap && sortedList[j-gap] > temp; j -= gap {
				sortedList[j] = sortedList[j-gap]
			}
			sortedList[j] = temp
		}
		if gap == 1 {
			gap = 0
		}
	}
	return sortedList
}
