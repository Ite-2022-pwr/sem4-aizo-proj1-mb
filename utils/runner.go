package utils

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"log"
	"math"
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
			floatList := make([]float64, len(ifh.GetFloat64List()))
			copy(floatList, ifh.GetShuffledFloat64List())
			runSort(method, ssgap, qspivot, i, floatList, ofh)
		}
	}
	ofh.WriteToFile()
}

func runSort[T constraints.Ordered](method string, ssgap, qspivot, rep int, list []T, ofh fileHandler.OutputFileHandler) {
	var timeTracked int64
	fmt.Print("pierwsze 10% lub 50 elementów listy przed sortowaniem: [")
	firstTenPercent := int(math.Ceil(float64(len(list)) * 0.1))
	if firstTenPercent > 50 {
		firstTenPercent = 50
	}
	for i := 0; i < firstTenPercent; i++ {
		fmt.Print(list[i], ", ")
	}
	fmt.Print("...]\n")
	fmt.Println()
	listAfterSort := make([]T, len(list))
	switch method {
	case "q":
		listAfterSort, timeTracked = mySort.QuickSort(list, 0, len(list)-1, qspivot)
	case "i":
		listAfterSort, timeTracked = mySort.InsertionSort(list)
	case "h":
		listAfterSort, timeTracked = mySort.HeapSort(list)
	case "s":
		listAfterSort, timeTracked = mySort.ShellSort(list, len(list), ssgap)
	default:
		listAfterSort, timeTracked = mySort.QuickSort(list, 0, len(list)-1, qspivot)
	}
	log.Printf("Posortowano poprawnie: %t", CheckSortedList(listAfterSort))
	ofh.AddResult(timeTracked, len(list), rep)
}

func RunSortFromFile(typeChosen int, sortType string, qsPivot, ssGap, reps int, ifh fileHandler.InputFileHandler, ofh fileHandler.OutputFileHandler) {
	switch typeChosen {
	case 0:
		for i := 0; i < reps; i++ {
			list := make([]int, len(ifh.GetIntList()))
			copy(list, ifh.GetIntList())
			runSort(sortType, ssGap, qsPivot, i, list, ofh)
		}
	case 1:
		for i := 0; i < reps; i++ {
			list := make([]float64, len(ifh.GetFloat64List()))
			copy(list, ifh.GetFloat64List())
			runSort(sortType, ssGap, qsPivot, i, list, ofh)
		}
	case 2:
		for i := 0; i < reps; i++ {
			list := make([]float32, len(ifh.GetFloat32List()))
			copy(list, ifh.GetFloat32List())
			runSort(sortType, ssGap, qsPivot, i, list, ofh)
		}
	case 3:
		for i := 0; i < reps; i++ {
			list := make([]int32, len(ifh.GetInt32List()))
			copy(list, ifh.GetInt32List())
			runSort(sortType, ssGap, qsPivot, i, list, ofh)
		}
	case 4:
		for i := 0; i < reps; i++ {
			list := make([]int64, len(ifh.GetInt64List()))
			copy(list, ifh.GetInt64List())
			runSort(sortType, ssGap, qsPivot, i, list, ofh)
		}
	case 5:
		for i := 0; i < reps; i++ {
			list := make([]string, len(ifh.GetStringList()))
			copy(list, ifh.GetStringList())
			runSort(sortType, ssGap, qsPivot, i, list, ofh)
		}
	default:
		for i := 0; i < reps; i++ {
			list := make([]int, len(ifh.GetIntList()))
			copy(list, ifh.GetIntList())
			runSort(sortType, ssGap, qsPivot, i, list, ofh)
		}
	}
}

func RunSortOfType(typeChosen int, sortType string, qsPivot, ssGap, reps, lines int, ofh fileHandler.OutputFileHandler) {
	switch typeChosen {
	case 0:
		list := GenerateRandomIntList(lines)
		for i := 0; i < reps; i++ {
			runSort(sortType, ssGap, qsPivot, i, list, ofh)
		}
	case 1:
		list := GenerateRandomFloat64List(lines)
		for i := 0; i < reps; i++ {
			runSort(sortType, ssGap, qsPivot, i, list, ofh)
		}
	case 2:
		list := GenerateRandomFloat32List(lines)
		for i := 0; i < reps; i++ {
			runSort(sortType, ssGap, qsPivot, i, list, ofh)
		}
	case 3:
		list := GenerateRandomInt32List(lines)
		for i := 0; i < reps; i++ {
			runSort(sortType, ssGap, qsPivot, i, list, ofh)
		}
	case 4:
		list := GenerateRandomInt64List(lines)
		for i := 0; i < reps; i++ {
			runSort(sortType, ssGap, qsPivot, i, list, ofh)
		}
	case 5:
		list := GenerateRandomStringList(lines)
		for i := 0; i < reps; i++ {
			runSort(sortType, ssGap, qsPivot, i, list, ofh)
		}
	default:
		list := GenerateRandomIntList(lines)
		for i := 0; i < reps; i++ {
			runSort(sortType, ssGap, qsPivot, i, list, ofh)
		}
	}
}
