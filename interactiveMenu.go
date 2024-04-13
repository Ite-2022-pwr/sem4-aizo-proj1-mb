package main

import (
	"bufio"
	"fmt"
	"os"
	"projekt1/fileHandler"
	"projekt1/utils"
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
	intTypeChoosen, _ := strconv.ParseInt(typeChosen, 10, 64)
	intQsPivot, _ := strconv.ParseInt(qsPivot, 10, 64)
	intShellGap, _ := strconv.ParseInt(shellGap, 10, 64)
	intLines, _ := strconv.ParseInt(lines, 10, 64)
	intRepeats, _ = strconv.ParseInt(repeats, 10, 64)

	if dataSource == "0" {
		ifh = fileHandler.InputFileHandler{FileName: filename}
		ifh.ReadFile()
		utils.RunSortFromFile(int(intTypeChoosen), sortingAlgorithm, int(intQsPivot), int(intShellGap), int(intRepeats), ifh, ofh)
	} else {
		utils.RunSortOfType(int(intTypeChoosen), sortingAlgorithm, int(intQsPivot), int(intShellGap), int(intRepeats), int(intLines), ofh)
	}

	ofh.WriteToFile()

}
