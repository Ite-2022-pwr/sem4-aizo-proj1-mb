package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"projekt1/fileHandler"
	"projekt1/mySort"
)

func Run(input, output, method string, ssgap, qspivot, vartype, reps, genlen int, newinput bool) {
	if newinput {
		GenerateInputToFile(genlen, vartype, input)
	}
	ifh := fileHandler.InputFileHandler{FileName: input}
	ifh.ReadFile()
	ofh := fileHandler.OutputFileHandler{FileName: output}
	ofh.InitializeSlices(reps)
	if vartype == 0 {
		for i := 0; i < reps; i++ {
			intList := make([]int, len(ifh.GetIntList()))
			copy(intList, ifh.GetShuffledIntList())
			runSort(method, ssgap, qspivot, i, intList, ofh)
		}
	} else if vartype == 1 {
		for i := 0; i < reps; i++ {
			floatList := make([]float64, len(ifh.GetFloatList()))
			copy(floatList, ifh.GetShuffledFloatList())
			fmt.Println(floatList)
			runSort(method, ssgap, qspivot, i, floatList, ofh)
		}
	}
	ofh.WriteToFile()
}

func runSort[T constraints.Ordered](method string, ssgap, qspivot, rep int, list []T, ofh fileHandler.OutputFileHandler) {
	var timeTracked int64
	var sortedList []T
	switch method {
	case "q":
		sortedList, timeTracked = mySort.QuickSort(list, 0, len(list)-1, qspivot)
	case "i":
		sortedList, timeTracked = mySort.InsertionSort(list)
	case "h":
		sortedList, timeTracked = mySort.HeapSort(list)
	case "s":
		sortedList, timeTracked = mySort.ShellSort(list, len(list), ssgap)
	default:
		sortedList, timeTracked = mySort.QuickSort(list, 0, len(list)-1, qspivot)
	}
	ofh.AddResult(timeTracked, len(list), rep)
	fmt.Println(sortedList)
}
