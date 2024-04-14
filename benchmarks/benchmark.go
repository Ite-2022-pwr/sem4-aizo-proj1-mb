package benchmarks

import (
	"fmt"
	"projekt1/fileHandler"
	"projekt1/utils"
)

func SingleConfigBenchmark(typeChosen, pivotChosen, ssGapChosen, repsChosen int, sortingMethodChosen, caseTypeChosen string) {
	typeName := utils.TypeNameChooser(typeChosen)
	pivotChosenName := utils.PivotNameChooser(pivotChosen)
	ssGapChosenName := utils.SsGapNameChooser(ssGapChosen)
	sortingMethodChosenName := utils.SortingMethodNameChooser(sortingMethodChosen)
	if sortingMethodChosen == "q" {
		for i := 1; i <= 10; i++ {
			ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("%s%s-%d.txt", caseTypeChosen, typeName, i)}
			ifh.ReadFile()
			ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("%s%sPiv%s%s-%d.csv", sortingMethodChosenName, caseTypeChosen, pivotChosenName, typeName, i)}
			ofh.InitializeSlices(repsChosen)
			utils.RunSortFromFile(typeChosen, sortingMethodChosen, pivotChosen, ssGapChosen, repsChosen, ifh, ofh)
			ofh.WriteToFile()
		}
	} else if sortingMethodChosen == "s" {
		for i := 1; i <= 10; i++ {
			ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("%s%s-%d.txt", caseTypeChosen, typeName, i)}
			ifh.ReadFile()
			ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("%s%s%s%s-%d.csv", sortingMethodChosenName, ssGapChosenName, caseTypeChosen, typeName, i)}
			ofh.InitializeSlices(repsChosen)
			utils.RunSortFromFile(typeChosen, sortingMethodChosen, pivotChosen, ssGapChosen, repsChosen, ifh, ofh)
			ofh.WriteToFile()
		}
	} else {
		for i := 1; i <= 10; i++ {
			ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("%s%s-%d.txt", caseTypeChosen, typeName, i)}
			ifh.ReadFile()
			ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("%s%s%s-%d.csv", sortingMethodChosenName, caseTypeChosen, typeName, i)}
			ofh.InitializeSlices(repsChosen)
			utils.RunSortFromFile(typeChosen, sortingMethodChosen, pivotChosen, ssGapChosen, repsChosen, ifh, ofh)
			ofh.WriteToFile()
		}
	}

}
