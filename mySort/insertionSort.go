package mySort

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"projekt1/timeTrack"
	"time"
)

func InsertionSort[T constraints.Ordered](list []T) []T {
	name := fmt.Sprintf("Insertion sort, type: %T", *new(T))
	defer timeTrack.TimeTrack(time.Now(), name)
	for i := 0; i < len(list); i++ {
		key := list[i]
		j := i - 1
		//roundTime := time.Now()
		for j >= 0 && list[j] > key {
			list[j+1] = list[j]
			j = j - 1
		}
		list[j+1] = key
		//elapsed := time.Since(roundTime)
		//log.Printf("Round %d took %s", i, elapsed)
	}
	return list
}
