package mySort

import (
	"golang.org/x/exp/constraints"
	"math/rand"
	"projekt1/timeTrack"
	"time"
)

func partitonHigh[T constraints.Ordered](list []T, low, high int) int {
	pivot := list[high]
	i := low - 1
	for j := low; j < high; j++ {
		if list[j] <= pivot {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	i++
	list[i], list[high] = list[high], list[i]
	return i
}

func partitonLow[T constraints.Ordered](list []T, low, high int) int {
	pivot := list[low]
	i := high
	for j := high; j > low; j-- {
		if list[j] >= pivot {
			list[i], list[j] = list[j], list[i]
			i--
		}
	}
	list[i], list[low] = list[low], list[i]
	return i
}

func partitonMid[T constraints.Ordered](list []T, low, high int) int {
	middle := low + ((high - low) / 2)
	list[high], list[middle] = list[middle], list[high]
	pivot := list[high]
	i := low - 1
	for j := low; j < high; j++ {
		if list[j] <= pivot {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	i++
	list[i], list[high] = list[high], list[i]
	return i
}

func partitonRand[T constraints.Ordered](list []T, low, high int) int {
	randomIndex := low + (rand.Int() % (high - low))
	list[high], list[randomIndex] = list[randomIndex], list[high]
	pivot := list[high]
	i := low - 1
	for j := low; j < high; j++ {
		if list[j] <= pivot {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	i++
	list[i], list[high] = list[high], list[i]
	return i
}

func quickSortPartitionHigh[T constraints.Ordered](list []T, low, high int) []T {
	if low == 0 && high == len(list)-1 {
		defer timeTrack.TimeTrack(time.Now(), "Quick sort, high partition")
	}
	if low >= high || low < 0 {
		return list
	}
	pivot := partitonHigh(list, low, high)

	quickSortPartitionHigh(list, low, pivot-1)
	quickSortPartitionHigh(list, pivot+1, high)
	return list
}

func quickSortPartitionLow[T constraints.Ordered](list []T, low, high int) []T {
	if low == 0 && high == len(list)-1 {
		defer timeTrack.TimeTrack(time.Now(), "Quick sort, low partition")
	}
	if low >= high || low < 0 {
		return list
	}
	pivot := partitonLow(list, low, high)

	quickSortPartitionLow(list, low, pivot-1)
	quickSortPartitionLow(list, pivot+1, high)
	return list
}

func quickSortPartitionMid[T constraints.Ordered](list []T, low, high int) []T {
	if low == 0 && high == len(list)-1 {
		defer timeTrack.TimeTrack(time.Now(), "Quick sort, mid partition")
	}
	if low >= high || low < 0 {
		return list
	}
	pivot := partitonMid(list, low, high)

	quickSortPartitionMid(list, low, pivot-1)
	quickSortPartitionMid(list, pivot+1, high)
	return list
}

func quickSortPartitionRand[T constraints.Ordered](list []T, low, high int) []T {
	if low == 0 && high == len(list)-1 {
		defer timeTrack.TimeTrack(time.Now(), "Quick sort, rand partition")
	}
	if low >= high || low < 0 {
		return list
	}
	pivot := partitonRand(list, low, high)

	quickSortPartitionRand(list, low, pivot-1)
	quickSortPartitionRand(list, pivot+1, high)
	return list
}

func QuickSort[T constraints.Ordered](list []T, low, high, partitonType int) []T {
	switch partitonType {
	case 0:
		return quickSortPartitionHigh(list, low, high)
	case 1:
		return quickSortPartitionLow(list, low, high)
	case 2:
		return quickSortPartitionMid(list, low, high)
	case 3:
		return quickSortPartitionRand(list, low, high)
	default:
		return quickSortPartitionHigh(list, low, high)
	}
}
