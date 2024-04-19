package runner

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"log"
	"math"
	"projekt1/fileHandler"
	"projekt1/mySort"
	"projekt1/utils"
	"time"
)

func Run(method, typeChosen, ssGap, qsPivot, reps int, list []any) (lengths []int, times []int64) {
	lengths = make([]int, reps)
	for i := 0; i < reps; i++ {
		lengths[i] = len(list)
	}
	times = make([]int64, reps)
	switch typeChosen {
	case 0:
		tempList := utils.CastToIntSlice(list)
		_, times = repeatSort(method, ssGap, qsPivot, reps, tempList)
	case 1:
		tempList := utils.CastToFloat64Slice(list)
		_, times = repeatSort(method, ssGap, qsPivot, reps, tempList)
	case 2:
		tempList := utils.CastToFloat32Slice(list)
		_, times = repeatSort(method, ssGap, qsPivot, reps, tempList)
	case 3:
		tempList := utils.CastToInt32Slice(list)
		_, times = repeatSort(method, ssGap, qsPivot, reps, tempList)
	case 4:
		tempList := utils.CastToInt64Slice(list)
		_, times = repeatSort(method, ssGap, qsPivot, reps, tempList)
	case 5:
		tempList := utils.CastToStringSlice(list)
		_, times = repeatSort(method, ssGap, qsPivot, reps, tempList)
	default:
		tempList := utils.CastToIntSlice(list)
		_, times = repeatSort(method, ssGap, qsPivot, reps, tempList)
	}
	return lengths, times
}

func RunFromFlags(inputFilename, outputFilename string, method, typeChosen, ssGap, qsPivot, reps, genLength, genStrat int, gen bool) {
	var list []any
	if gen {
		list = utils.GenerateList(typeChosen, genStrat, genLength)
	} else {
		ifh := fileHandler.InputFileHandler{FileName: inputFilename}
		ifh.ReadFile()
		list = ifh.GetListOfType(typeChosen)
	}
	lengths, times := Run(method, typeChosen, ssGap, qsPivot, reps, list)
	ofh := fileHandler.OutputFileHandler{FileName: outputFilename}
	ofh.InitializeSlices(reps)
	ofh.AddResultLists(lengths, times)
	ofh.WriteToFile()
}

func repeatSort[T constraints.Ordered](method, ssGap, qsPivot, reps int, list []T) ([]T, []int64) {
	var timesTracked []int64
	timesTracked = make([]int64, reps)
	var totalTime int64 = 0
	for i := 0; i < reps; i++ {
		tempList := make([]T, len(list))
		copy(tempList, list)
		_, timesTracked[i] = runSort(method, ssGap, qsPivot, tempList)
	}
	for _, trackedTime := range timesTracked {
		totalTime += trackedTime
	}
	avgTime := totalTime / int64(reps)
	avgTimeDuration := time.Duration(avgTime)
	log.Printf("Średni czas sortowania: %s", avgTimeDuration.String())
	return list, timesTracked
}

func runSort[T constraints.Ordered](method, ssgap, qspivot int, list []T) ([]T, int64) {
	var timeTracked int64
	tempStr := "pierwsze 10% lub 50 elementów listy przed sortowaniem: ["
	firstTenPercent := int(math.Min(math.Ceil(float64(len(list))*0.1), 50))
	for i := 0; i < firstTenPercent; i++ {
		tempStr += fmt.Sprintf("%v; ", list[i])
	}
	tempStr += fmt.Sprintf("...]\n")
	log.Println(tempStr)
	listAfterSort := make([]T, len(list))
	switch method {
	case 0:
		listAfterSort, timeTracked = mySort.QuickSort(list, 0, len(list)-1, qspivot)
	case 1:
		listAfterSort, timeTracked = mySort.HeapSort(list)
	case 2:
		listAfterSort, timeTracked = mySort.InsertionSort(list)
	case 3:
		listAfterSort, timeTracked = mySort.ShellSort(list, len(list), ssgap)
	default:
		listAfterSort, timeTracked = mySort.QuickSort(list, 0, len(list)-1, qspivot)
	}
	log.Printf("Posortowano poprawnie: %t", utils.CheckSortedList(listAfterSort))
	tempStr = "pierwsze 10% lub 50 elementów listy po sortowaniu: ["
	for i := 0; i < firstTenPercent; i++ {
		tempStr += fmt.Sprintf("%v; ", listAfterSort[i])
	}
	tempStr += fmt.Sprintf("...]\n")
	log.Println(tempStr)
	return list, timeTracked
}
