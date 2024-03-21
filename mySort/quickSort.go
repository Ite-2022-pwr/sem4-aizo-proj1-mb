package mySort

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math/rand"
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
	pivot := list[low+((high-low)/2)]
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

func partitonRand[T constraints.Ordered](list []T, low, high int) int {
	randomIndex := low + (rand.Int() % (high - low))
	list[high], list[randomIndex] = list[randomIndex], list[high]
	pivot := list[high]
	fmt.Println("pivot: ", pivot)
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

func QuickSort[T constraints.Ordered](list []T, low, high, partitonType int) []T {
	if low >= high || low < 0 {
		return list
	}
	var pivot int
	switch partitonType {
	case 0:
		pivot = partitonHigh(list, low, high)
	case 1:
		pivot = partitonLow(list, low, high)
	case 2:
		pivot = partitonMid(list, low, high)
	case 3:
		pivot = partitonRand(list, low, high)
	default:
		pivot = partitonHigh(list, low, high)
	}
	QuickSort(list, low, pivot-1, partitonType)
	QuickSort(list, pivot+1, high, partitonType)
	return list
}
