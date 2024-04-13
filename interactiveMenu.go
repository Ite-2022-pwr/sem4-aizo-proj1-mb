package main

import (
	"bufio"
	"fmt"
	"os"
	"projekt1/fileHandler"
	"strconv"
	"strings"
)

func Menu() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Wybierz typ zmiennej do testowania:")
	fmt.Println("0 - int")
	fmt.Println("1 - float64")
	fmt.Println("2 - float32")
	fmt.Println("3 - int32")
	fmt.Println("4 - int64")
	fmt.Println("5 - string")
	typeChosen, _ := reader.ReadString('\n')
	typeChosen, _, _ = strings.Cut(typeChosen, "\n")
	fmt.Println("Wybierz algorytm do testowania:")
	fmt.Println("h - Heap sort")
	fmt.Println("s - Shell sort")
	fmt.Println("q - Quick sort")
	fmt.Println("i - insertion sort")
	sortingAlgorithm, _ := reader.ReadString('\n')
	strings.Cut(sortingAlgorithm, "\n")
	sortingAlgorithm, _, _ = strings.Cut(sortingAlgorithm, "\n")
	var qsPivot string
	if sortingAlgorithm == "q" {
		fmt.Println("Wybierz typ pivotu:")
		fmt.Println("0 - ostatni element")
		fmt.Println("1 - pierwszy element")
		fmt.Println("2 - środkowy element")
		fmt.Println("3 - losowy element")
		qsPivot, _ = reader.ReadString('\n')
		strings.Cut(qsPivot, "\n")
		qsPivot, _, _ = strings.Cut(qsPivot, "\n")
	}
	var shellGap string
	if sortingAlgorithm == "s" {
		fmt.Println("Wybierz Typ odstępu:")
		fmt.Println("0 - Shell")
		fmt.Println("1 - Frank & Lazarus")
		shellGap, _ = reader.ReadString('\n')
		strings.Cut(shellGap, "\n")
		shellGap, _, _ = strings.Cut(shellGap, "\n")
	}
	fmt.Println("Wybierz sposób generowania danych:")
	fmt.Println("0 - z pliku")
	fmt.Println("1 - losowe")
	dataSource, _ := reader.ReadString('\n')
	strings.Cut(dataSource, "\n")
	dataSource, _, _ = strings.Cut(dataSource, "\n")
	var filename string
	if dataSource == "0" {
		fmt.Println("Podaj nazwę pliku:")
		filename, _ = reader.ReadString('\n')
		strings.Cut(filename, "\n")
		filename, _, _ = strings.Cut(filename, "\n")
	}
	var lines string
	if dataSource == "1" {
		fmt.Println("Podaj ilość danych:")
		lines, _ = reader.ReadString('\n')
		strings.Cut(lines, "\n")
		lines, _, _ = strings.Cut(lines, "\n")
	}
	fmt.Println("Podaj ilość powtórzeń:")
	repeats, _ := reader.ReadString('\n')
	strings.Cut(repeats, "\n")
	repeats, _, _ = strings.Cut(repeats, "\n")
	fmt.Println("podaj nazwę pliku wynikowego:")
	outputFilename, _ := reader.ReadString('\n')
	strings.Cut(outputFilename, "\n")
	outputFilename, _, _ = strings.Cut(outputFilename, "\n")

	var ifh fileHandler.InputFileHandler
	ofh := fileHandler.OutputFileHandler{FileName: outputFilename}
	intRepeats, _ := strconv.ParseInt(repeats, 10, 64)
	ofh.InitializeSlices(int(intRepeats))

	if dataSource == "0" {
		ifh = fileHandler.InputFileHandler{FileName: filename}
		ifh.ReadFile()
		runSortFromFile(typeChosen, sortingAlgorithm, qsPivot, shellGap, repeats, ifh, ofh)
	} else {
		runSortOfType(typeChosen, sortingAlgorithm, qsPivot, shellGap, repeats, lines, ofh)
	}

	ofh.WriteToFile()

}

func runSortFromFile(typeChosen, sortType, qsPivot, ssGap, reps string, ifh fileHandler.InputFileHandler, ofh fileHandler.OutputFileHandler) {
	intTypeChosen, _ := strconv.ParseInt(typeChosen, 10, 64)
	intSsGap, _ := strconv.ParseInt(ssGap, 10, 64)
	intQsPivot, _ := strconv.ParseInt(qsPivot, 10, 64)
	intReps, _ := strconv.ParseInt(reps, 10, 64)

	switch intTypeChosen {
	case 0:
		list := ifh.GetIntList()
		for i := 0; i < int(intReps); i++ {
			runSort(sortType, int(intSsGap), int(intQsPivot), i, list, ofh)
		}
	case 1:
		list := ifh.GetFloat64List()
		for i := 0; i < int(intReps); i++ {
			runSort(sortType, int(intSsGap), int(intQsPivot), i, list, ofh)
		}
	case 2:
		list := ifh.GetFloat32List()
		for i := 0; i < int(intReps); i++ {
			runSort(sortType, int(intSsGap), int(intQsPivot), i, list, ofh)
		}
	case 3:
		list := ifh.GetInt32List()
		for i := 0; i < int(intReps); i++ {
			runSort(sortType, int(intSsGap), int(intQsPivot), i, list, ofh)
		}
	case 4:
		list := ifh.GetInt64List()
		for i := 0; i < int(intReps); i++ {
			runSort(sortType, int(intSsGap), int(intQsPivot), i, list, ofh)
		}
	case 5:
		list := ifh.GetStringList()
		for i := 0; i < int(intReps); i++ {
			runSort(sortType, int(intSsGap), int(intQsPivot), i, list, ofh)
		}
	default:
		list := ifh.GetIntList()
		for i := 0; i < int(intReps); i++ {
			runSort(sortType, int(intSsGap), int(intQsPivot), i, list, ofh)
		}
	}
}

func runSortOfType(typeChosen, sortType, qsPivot, ssGap, reps, lines string, ofh fileHandler.OutputFileHandler) {
	intTypeChosen, _ := strconv.ParseInt(typeChosen, 10, 64)
	intSsGap, _ := strconv.ParseInt(ssGap, 10, 64)
	intQsPivot, _ := strconv.ParseInt(qsPivot, 10, 64)
	intReps, _ := strconv.ParseInt(reps, 10, 64)
	intLines, _ := strconv.ParseInt(lines, 10, 64)
	switch intTypeChosen {
	case 0:
		list := GenerateRandomIntList(int(intLines))
		for i := 0; i < int(intReps); i++ {
			runSort(sortType, int(intSsGap), int(intQsPivot), i, list, ofh)
		}
	case 1:
		list := GenerateRandomFloat64List(int(intLines))
		for i := 0; i < int(intReps); i++ {
			runSort(sortType, int(intSsGap), int(intQsPivot), i, list, ofh)
		}
	case 2:
		list := GenerateRandomFloat32List(int(intLines))
		for i := 0; i < int(intReps); i++ {
			runSort(sortType, int(intSsGap), int(intQsPivot), i, list, ofh)
		}
	case 3:
		list := GenerateRandomInt32List(int(intLines))
		for i := 0; i < int(intReps); i++ {
			runSort(sortType, int(intSsGap), int(intQsPivot), i, list, ofh)
		}
	case 4:
		list := GenerateRandomInt64List(int(intLines))
		for i := 0; i < int(intReps); i++ {
			runSort(sortType, int(intSsGap), int(intQsPivot), i, list, ofh)
		}
	case 5:
		list := GenerateRandomStringList(int(intLines))
		for i := 0; i < int(intReps); i++ {
			runSort(sortType, int(intSsGap), int(intQsPivot), i, list, ofh)
		}
	default:
		list := GenerateRandomIntList(int(intLines))
		for i := 0; i < int(intReps); i++ {
			runSort(sortType, int(intSsGap), int(intQsPivot), i, list, ofh)
		}
	}
}
