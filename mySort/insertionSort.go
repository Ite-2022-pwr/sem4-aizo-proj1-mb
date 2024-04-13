package mySort

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"log"
	"projekt1/timeTrack"
	"time"
)

func InsertionSort[T constraints.Ordered](list []T) (sortedList []T, timeTracked int64) {
	name := fmt.Sprintf("Insertion sort, na typie: %T", *new(T))
	startTime := time.Now()
	log.Printf("%s rozpoczęto o: %s", name, startTime)
	defer func() {
		timeTracked = timeTrack.TimeTrack(startTime, name)
	}()
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
	return list, 0
}
