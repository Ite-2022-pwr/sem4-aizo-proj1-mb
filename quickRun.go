package main

import (
	"log"
	"projekt1/fileHandler"
	"projekt1/mySort"
	"sort"
	"sync"
)

func QuickRun() {
	n := 100000
	//generateInput(n, 1, "floaty.txt")
	//generateInput(n, 0, "inty.txt")
	fhF := fileHandler.InputFileHandler{FileName: "floaty.txt"}
	fhF.ReadFile()

	fhI := fileHandler.InputFileHandler{FileName: "inty.txt"}
	fhI.ReadFile()
	floatlists := make([][]float64, 7)
	intLists := make([][]int, 7)
	var wg sync.WaitGroup
	for i := 0; i < 7; i++ {
		wg.Add(1)
		i := i
		go func() {
			floatlists[i] = make([]float64, n)
			intLists[i] = make([]int, n)
			copy(floatlists[i], fhF.GetFloat64List())
			copy(intLists[i], fhI.GetIntList())
			wg.Done()
		}()
	}
	wg.Wait()

	var sortFWG sync.WaitGroup
	sortFWG.Add(7)

	go func() {
		mySort.InsertionSort(floatlists[0])
		log.Printf("properly sorted %v", sort.Float64sAreSorted(floatlists[0]))
		sortFWG.Done()
	}()
	go func() {
		mySort.HeapSort(floatlists[1])
		log.Printf("properly sorted %v", sort.Float64sAreSorted(floatlists[1]))
		sortFWG.Done()
	}()
	go func() {
		mySort.ShellSort(floatlists[2], len(floatlists[2]), 1)
		log.Printf("properly sorted %v", sort.Float64sAreSorted(floatlists[2]))
		sortFWG.Done()
	}()
	go func() {
		mySort.QuickSort(floatlists[3], 0, len(floatlists[3])-1, 0)
		log.Printf("properly sorted %v", sort.Float64sAreSorted(floatlists[3]))
		sortFWG.Done()
	}()
	go func() {
		mySort.QuickSort(floatlists[4], 0, len(floatlists[4])-1, 1)
		log.Printf("properly sorted %v", sort.Float64sAreSorted(floatlists[4]))
		sortFWG.Done()
	}()
	go func() {
		mySort.QuickSort(floatlists[5], 0, len(floatlists[5])-1, 2)
		log.Printf("properly sorted %v", sort.Float64sAreSorted(floatlists[5]))
		sortFWG.Done()
	}()
	go func() {
		mySort.QuickSort(floatlists[6], 0, len(floatlists[6])-1, 3)
		log.Printf("properly sorted %v", sort.Float64sAreSorted(floatlists[6]))
		sortFWG.Done()
	}()
	sortFWG.Wait()

	var sortIWG sync.WaitGroup
	sortIWG.Add(7)

	go func() {
		mySort.InsertionSort(intLists[0])
		log.Printf("properly sorted %v", sort.IntsAreSorted(intLists[0]))
		sortIWG.Done()
	}()
	go func() {
		mySort.HeapSort(intLists[1])
		log.Printf("properly sorted %v", sort.IntsAreSorted(intLists[1]))
		sortIWG.Done()
	}()
	go func() {
		mySort.ShellSort(intLists[2], len(intLists[2]), 1)
		log.Printf("properly sorted %v", sort.IntsAreSorted(intLists[2]))
		sortIWG.Done()
	}()
	go func() {
		mySort.QuickSort(intLists[3], 0, len(intLists[3])-1, 0)
		log.Printf("properly sorted %v", sort.IntsAreSorted(intLists[3]))
		sortIWG.Done()
	}()
	go func() {
		mySort.QuickSort(intLists[4], 0, len(intLists[4])-1, 1)
		log.Printf("properly sorted %v", sort.IntsAreSorted(intLists[4]))
		sortIWG.Done()
	}()
	go func() {
		mySort.QuickSort(intLists[5], 0, len(intLists[5])-1, 2)
		log.Printf("properly sorted %v", sort.IntsAreSorted(intLists[5]))
		sortIWG.Done()
	}()
	go func() {
		mySort.QuickSort(intLists[6], 0, len(intLists[6])-1, 3)
		log.Printf("properly sorted %v", sort.IntsAreSorted(intLists[6]))
		sortIWG.Done()
	}()
	sortIWG.Wait()

}
