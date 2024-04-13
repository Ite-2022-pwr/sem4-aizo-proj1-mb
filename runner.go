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
	ofh := fileHandler.OutputFileHandler{FileName: output}
	ofh.InitializeSlices(reps)
	if vartype == 0 {
		for i := 0; i < reps; i++ {
			intList := make([]int, len(ifh.GetIntList()))
			copy(intList, ifh.GetIntList())
			runSort(method, ssgap, qspivot, i, intList, ofh)
		}
	} else if vartype == 1 {
		for i := 0; i < reps; i++ {
			floatList := make([]float64, len(ifh.GetFloatList()))
			copy(floatList, ifh.GetFloatList())
			runSort(method, ssgap, qspivot, i, floatList, ofh)
		}
	}
	ofh.WriteToFile()
}

func runSort[T constraints.Ordered](method string, ssgap, qspivot, rep int, list []T, ofh fileHandler.OutputFileHandler) {
	var timeTracked int64
	switch method {
	case "q":
		_, timeTracked = mySort.QuickSort(list, 0, len(list)-1, qspivot)
	case "i":
		_, timeTracked = mySort.InsertionSort(list)
	case "h":
		_, timeTracked = mySort.HeapSort(list)
	case "s":
		_, timeTracked = mySort.ShellSort(list, len(list), ssgap)
	default:
		_, timeTracked = mySort.QuickSort(list, 0, len(list)-1, qspivot)
	}
	ofh.AddResult(timeTracked, len(list), rep)
}
