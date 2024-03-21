package main

import (
	"projekt1/fileHandler"
	"projekt1/mySort"
	"sync"
)

func main() {
	generateInput(1000000, 1, "floaty.txt")
	fh := fileHandler.FileHandler{FileName: "floaty.txt"}
	fh.ReadFile()
	floatListaI := make([]float64, len(fh.GetFloatList()))
	floatListaH := make([]float64, len(fh.GetFloatList()))
	floatListaS := make([]float64, len(fh.GetFloatList()))
	floatListaQH := make([]float64, len(fh.GetFloatList()))
	floatListaQL := make([]float64, len(fh.GetFloatList()))
	floatListaQM := make([]float64, len(fh.GetFloatList()))
	floatListaQR := make([]float64, len(fh.GetFloatList()))
	copy(floatListaI, fh.GetFloatList())
	copy(floatListaH, fh.GetFloatList())
	copy(floatListaS, fh.GetFloatList())
	copy(floatListaQH, fh.GetFloatList())
	copy(floatListaQL, fh.GetFloatList())
	copy(floatListaQM, fh.GetFloatList())
	copy(floatListaQR, fh.GetFloatList())

	var wg sync.WaitGroup
	wg.Add(7)

	go func() {
		mySort.InsertionSort(floatListaI)
		wg.Done()
	}()
	go func() {
		mySort.HeapSort(floatListaH)
		wg.Done()
	}()
	go func() {
		mySort.ShellSort(floatListaS, len(floatListaS), 1)
		wg.Done()
	}()
	go func() {
		mySort.QuickSort(floatListaQH, 0, len(floatListaQH)-1, 0)
		wg.Done()
	}()
	go func() {
		mySort.QuickSort(floatListaQL, 0, len(floatListaQL)-1, 1)
		wg.Done()
	}()
	go func() {
		mySort.QuickSort(floatListaQM, 0, len(floatListaQM)-1, 2)
		wg.Done()
	}()
	go func() {
		mySort.QuickSort(floatListaQR, 0, len(floatListaQR)-1, 3)
		wg.Done()
	}()

	wg.Wait()
}
