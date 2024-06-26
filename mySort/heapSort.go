package mySort

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"log"
	"projekt1/timeTrack"
	"time"
)

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

func HeapSort[T constraints.Ordered](list []T) (sortedList []T, timeTracked int64) {
	name := fmt.Sprintf("Heap sort, na typie: %T", *new(T))
	startTime := time.Now()
	log.Printf("%s rozpoczęto o: %s", name, startTime)
	defer func() {
		timeTracked = timeTrack.TimeTrack(startTime, name)
	}()
	n := len(list)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(list, n, i)
	}
	for i := n - 1; i > 0; i-- {
		list[0], list[i] = list[i], list[0]
		heapify(list, i, 0)
	}
	return list, 0
}
