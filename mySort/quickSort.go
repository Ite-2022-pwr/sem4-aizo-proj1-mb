package mySort

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"log"
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

func quickSortPartitionHigh[T constraints.Ordered](list []T, low, high int) ([]T, int64) {
	var timeTracked int64
	if low == 0 && high == len(list)-1 {
		name := fmt.Sprintf("Quick sort, high partition, type: %T", *new(T))
		startTime := time.Now()
		log.Printf("%s started at: %s", name, startTime)
		defer func() {
			timeTracked = timeTrack.TimeTrack(startTime, name)
		}()
	}
	if low >= high || low < 0 {
		return list, timeTracked
	}
	pivot := partitonHigh(list, low, high)

	quickSortPartitionHigh(list, low, pivot-1)
	quickSortPartitionHigh(list, pivot+1, high)
	return list, timeTracked
}

func quickSortPartitionLow[T constraints.Ordered](list []T, low, high int) ([]T, int64) {
	var timeTracked int64
	if low == 0 && high == len(list)-1 {
		name := fmt.Sprintf("Quick sort, low partition, type: %T", *new(T))
		startTime := time.Now()
		log.Printf("%s started at: %s", name, startTime)
		defer func() {
			timeTracked = timeTrack.TimeTrack(startTime, name)
		}()
	}
	if low >= high || low < 0 {
		return list, timeTracked
	}
	pivot := partitonLow(list, low, high)

	quickSortPartitionLow(list, low, pivot-1)
	quickSortPartitionLow(list, pivot+1, high)
	return list, timeTracked
}

func quickSortPartitionMid[T constraints.Ordered](list []T, low, high int) ([]T, int64) {
	var timeTracked int64
	if low == 0 && high == len(list)-1 {
		name := fmt.Sprintf("Quick sort, mid partition, type: %T", *new(T))
		startTime := time.Now()
		log.Printf("%s started at: %s", name, startTime)
		defer func() {
			timeTracked = timeTrack.TimeTrack(startTime, name)
		}()
	}
	if low >= high || low < 0 {
		return list, timeTracked
	}
	pivot := partitonMid(list, low, high)

	quickSortPartitionMid(list, low, pivot-1)
	quickSortPartitionMid(list, pivot+1, high)
	return list, timeTracked
}

func quickSortPartitionRand[T constraints.Ordered](list []T, low, high int) ([]T, int64) {
	var timeTracked int64
	if low == 0 && high == len(list)-1 {
		name := fmt.Sprintf("Quick sort, rand partition, type: %T", *new(T))
		startTime := time.Now()
		log.Printf("%s started at: %s", name, startTime)
		defer func() {
			timeTracked = timeTrack.TimeTrack(startTime, name)
		}()
	}
	if low >= high || low < 0 {
		return list, timeTracked
	}
	pivot := partitonRand(list, low, high)

	quickSortPartitionRand(list, low, pivot-1)
	quickSortPartitionRand(list, pivot+1, high)
	return list, timeTracked
}

func QuickSort[T constraints.Ordered](list []T, low, high, partitonType int) ([]T, int64) {
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
