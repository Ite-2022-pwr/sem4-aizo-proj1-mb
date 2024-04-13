package main

import (
	"golang.org/x/exp/constraints"
	"projekt1/fileHandler"
	"projekt1/mySort"
)

func Run(input, output, method string, ssgap, qspivot, vartype, reps, genlen int, newinput bool) {
	if newinput {
		generateInputToFile(genlen, vartype, input)
	}
	ifh := fileHandler.InputFileHandler{FileName: input}
	ifh.ReadFile()
	if vartype == 0 {
		for i := 0; i < reps; i++ {
			intList := make([]int, len(ifh.GetIntList()))
			copy(intList, ifh.GetIntList())
			runSort(method, ssgap, qspivot, intList)
		}
	} else if vartype == 1 {
		for i := 0; i < reps; i++ {
			floatList := make([]float64, len(ifh.GetFloatList()))
			copy(floatList, ifh.GetFloatList())
			runSort(method, ssgap, qspivot, floatList)
		}
	}

}

func runSort[T constraints.Ordered](method string, ssgap, qspivot int, list []T) {
	switch method {
	case "q":
		mySort.QuickSort(list, 0, len(list)-1, qspivot)
	case "i":
		mySort.InsertionSort(list)
	case "h":
		mySort.HeapSort(list)
	case "s":
		mySort.ShellSort(list, len(list), ssgap)
	default:
		mySort.QuickSort(list, 0, len(list)-1, qspivot)
	}
}
