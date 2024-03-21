package mySort

import (
	"golang.org/x/exp/constraints"
	"math"
	"projekt1/timeTrack"
	"time"
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
	defer timeTrack.TimeTrack(time.Now(), "Shell sort")
	k, gap := 1, n
	for gap >= 1 {
		gap = calculateGap(n, k, gapType)
		k++
		for i := gap; i < n; i++ {
			temp := list[i]
			j := i
			for j = j; j >= gap && list[j-gap] > temp; j -= gap {
				list[j] = list[j-gap]
			}
			list[j] = temp
		}
		if gap == 1 {
			gap = 0
		}
	}
	return list
}
