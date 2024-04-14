package benchmarks

import (
	"fmt"
	"log"
	"projekt1/fileHandler"
	"projekt1/utils"
)

func SingleConfigBenchmark(typeChosen, pivotChosen, ssGapChosen, repsChosen int, sortingMethodChosen, caseTypeChosen string) {
	typeName := utils.TypeNameChooser(typeChosen)
	pivotChosenName := utils.PivotNameChooser(pivotChosen)
	ssGapChosenName := utils.SsGapNameChooser(ssGapChosen)
	sortingMethodChosenName := utils.SortingMethodNameChooser(sortingMethodChosen)
	if sortingMethodChosen == "q" {
		for i := 0; i <= 7; i++ {
			ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("%s%s-%d.txt", caseTypeChosen, typeName, i)}
			ifh.ReadFile()
			log.Printf("Sortowanie %s z pivotem %s dla typu %s, na przypadku tablicy %s,  wielkości %d linii", sortingMethodChosenName, pivotChosenName, typeName, caseTypeChosen, ifh.Lines)
			ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("%s%sPiv%s%s-%d.csv", sortingMethodChosenName, caseTypeChosen, pivotChosenName, typeName, i)}
			ofh.InitializeSlices(repsChosen)
			utils.RunSortFromFile(typeChosen, sortingMethodChosen, pivotChosen, ssGapChosen, repsChosen, ifh, ofh)
			ofh.WriteToFile()
		}
	} else if sortingMethodChosen == "s" {
		for i := 0; i <= 7; i++ {
			ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("%s%s-%d.txt", caseTypeChosen, typeName, i)}
			ifh.ReadFile()
			log.Printf("Sortowanie %s z przerwą %s dla typu %s, na przypadku tablicy %s,  wielkości %d linii", sortingMethodChosenName, ssGapChosenName, typeName, caseTypeChosen, ifh.Lines)
			ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("%s%s%s%s-%d.csv", sortingMethodChosenName, ssGapChosenName, caseTypeChosen, typeName, i)}
			ofh.InitializeSlices(repsChosen)
			utils.RunSortFromFile(typeChosen, sortingMethodChosen, pivotChosen, ssGapChosen, repsChosen, ifh, ofh)
			ofh.WriteToFile()
		}
	} else {
		for i := 0; i <= 7; i++ {
			ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("%s%s-%d.txt", caseTypeChosen, typeName, i)}
			ifh.ReadFile()
			log.Printf("Sortowanie %s dla typu %s, na przypadku tablicy %s, wielkości %d linii", sortingMethodChosenName, typeName, caseTypeChosen, ifh.Lines)
			ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("%s%s%s-%d.csv", sortingMethodChosenName, caseTypeChosen, typeName, i)}
			ofh.InitializeSlices(repsChosen)
			utils.RunSortFromFile(typeChosen, sortingMethodChosen, pivotChosen, ssGapChosen, repsChosen, ifh, ofh)
			ofh.WriteToFile()
		}
	}

}
