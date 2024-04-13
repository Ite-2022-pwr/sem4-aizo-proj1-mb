package benchmarks

import (
	"fmt"
	"projekt1/fileHandler"
	"projekt1/utils"
)

func QuicksortBenchmark() {
	quicksortBenchmarkInt()
}

func quicksortBenchmarkInt() {
	quicksortBenchmarkIntSorted()
	quicksortBenchmarkIntReversed()
	quicksortBenchmarkIntRandom()
	quicksortBenchmarkIntThird()
}

func quicksortBenchmarkIntSorted() {
	for i := 1; i <= 10; i++ {
		ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("sortedInt-%d.txt", i)}
		ifh.ReadFile()
		ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("quicksortSortedPivHighInt-%d.csv", i)}
		ofh.InitializeSlices(25)
		utils.RunSortFromFile(0, "q", 0, 0, 25, ifh, ofh)
		ofh.WriteToFile()
	}
	for i := 1; i <= 10; i++ {
		ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("sortedInt-%d.txt", i)}
		ifh.ReadFile()
		ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("quicksortSortedPivLowInt-%d.csv", i)}
		ofh.InitializeSlices(25)
		utils.RunSortFromFile(0, "q", 1, 0, 25, ifh, ofh)
		ofh.WriteToFile()
	}
	for i := 1; i <= 10; i++ {
		ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("sortedInt-%d.txt", i)}
		ifh.ReadFile()
		ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("quicksortSortedPivMedInt-%d.csv", i)}
		ofh.InitializeSlices(25)
		utils.RunSortFromFile(0, "q", 2, 0, 25, ifh, ofh)
		ofh.WriteToFile()
	}
	for i := 1; i <= 10; i++ {
		ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("sortedInt-%d.txt", i)}
		ifh.ReadFile()
		ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("quicksortSortedPivRandInt-%d.csv", i)}
		ofh.InitializeSlices(25)
		utils.RunSortFromFile(0, "q", 3, 0, 25, ifh, ofh)
		ofh.WriteToFile()
	}
}

func quicksortBenchmarkIntReversed() {
	for i := 1; i <= 10; i++ {
		ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("reverseSortedInt-%d.txt", i)}
		ifh.ReadFile()
		ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("quicksortReversedPivHighInt-%d.csv", i)}
		ofh.InitializeSlices(25)
		utils.RunSortFromFile(0, "q", 0, 0, 25, ifh, ofh)
		ofh.WriteToFile()
	}
	for i := 1; i <= 10; i++ {
		ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("reverseSortedInt-%d.txt", i)}
		ifh.ReadFile()
		ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("quicksortReversedPivLowInt-%d.csv", i)}
		ofh.InitializeSlices(25)
		utils.RunSortFromFile(0, "q", 1, 0, 25, ifh, ofh)
		ofh.WriteToFile()
	}
	for i := 1; i <= 10; i++ {
		ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("reverseSortedInt-%d.txt", i)}
		ifh.ReadFile()
		ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("quicksortReversedPivMedInt-%d.csv", i)}
		utils.RunSortFromFile(0, "q", 2, 0, 25, ifh, ofh)
		ofh.WriteToFile()
	}
	for i := 1; i <= 10; i++ {
		ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("reverseSortedInt-%d.txt", i)}
		ifh.ReadFile()
		ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("quicksortReversedPivRandInt-%d.csv", i)}
		ofh.InitializeSlices(25)
		utils.RunSortFromFile(0, "q", 3, 0, 25, ifh, ofh)
		ofh.WriteToFile()
	}
}

func quicksortBenchmarkIntRandom() {
	for i := 1; i <= 10; i++ {
		ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("randomInt-%d.txt", i)}
		ifh.ReadFile()
		ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("quicksortRandomPivHighInt-%d.csv", i)}
		ofh.InitializeSlices(25)
		utils.RunSortFromFile(0, "q", 0, 0, 25, ifh, ofh)
		ofh.WriteToFile()
	}
	for i := 1; i <= 10; i++ {
		ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("randomInt-%d.txt", i)}
		ifh.ReadFile()
		ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("quicksortRandomPivLowInt-%d.csv", i)}
		ofh.InitializeSlices(25)
		utils.RunSortFromFile(0, "q", 1, 0, 25, ifh, ofh)
		ofh.WriteToFile()
	}
	for i := 1; i <= 10; i++ {
		ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("randomInt-%d.txt", i)}
		ifh.ReadFile()
		ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("quicksortRandomPivMedInt-%d.csv", i)}
		ofh.InitializeSlices(25)
		utils.RunSortFromFile(0, "q", 2, 0, 25, ifh, ofh)
		ofh.WriteToFile()
	}
	for i := 1; i <= 10; i++ {
		ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("randomInt-%d.txt", i)}
		ifh.ReadFile()
		ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("quicksortRandomPivRandInt-%d.csv", i)}
		ofh.InitializeSlices(25)
		utils.RunSortFromFile(0, "q", 3, 0, 25, ifh, ofh)
		ofh.WriteToFile()
	}
}

func quicksortBenchmarkIntThird() {
	for i := 1; i <= 10; i++ {
		ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("thirdSortedInt-%d.txt", i)}
		ifh.ReadFile()
		ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("quicksortThirdPivHighInt-%d.csv", i)}
		ofh.InitializeSlices(25)
		utils.RunSortFromFile(0, "q", 0, 0, 25, ifh, ofh)
		ofh.WriteToFile()
	}
	for i := 1; i <= 10; i++ {
		ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("thirdSortedInt-%d.txt", i)}
		ifh.ReadFile()
		ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("quicksortThirdPivLowInt-%d.csv", i)}
		ofh.InitializeSlices(25)
		utils.RunSortFromFile(0, "q", 1, 0, 25, ifh, ofh)
		ofh.WriteToFile()
	}
	for i := 1; i <= 10; i++ {
		ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("thirdSortedInt-%d.txt", i)}
		ifh.ReadFile()
		ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("quicksortThirdPivMedInt-%d.csv", i)}
		ofh.InitializeSlices(25)
		utils.RunSortFromFile(0, "q", 2, 0, 25, ifh, ofh)
		ofh.WriteToFile()
	}
	for i := 1; i <= 10; i++ {
		ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("thirdSortedInt-%d.txt", i)}
		ifh.ReadFile()
		ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("quicksortThirdPivRandInt-%d.csv", i)}
		ofh.InitializeSlices(25)
		utils.RunSortFromFile(0, "q", 3, 0, 25, ifh, ofh)
		ofh.WriteToFile()
	}
}
