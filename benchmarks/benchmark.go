package benchmarks

import (
	"fmt"
	"log"
	"projekt1/fileHandler"
	"projekt1/runner"
	"projekt1/utils"
)

func SingleConfigBenchmark(typeChosen, pivotChosen, ssGapChosen, repsChosen, sortingMethodChosen, caseTypeChosen int) {
	typeName := utils.TypeNameChooser(typeChosen)
	pivotChosenName := utils.PivotNameChooser(pivotChosen)
	ssGapChosenName := utils.SsGapNameChooser(ssGapChosen)
	sortingMethodChosenName := utils.SortingMethodNameChooser(sortingMethodChosen)
	if sortingMethodChosen == 0 {
		for i := 0; i < 7; i++ {
			ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("%s%s-%d.txt", caseTypeChosen, typeName, i)}
			ifh.ReadFile()
			log.Printf("Sortowanie %s z pivotem %s dla typu %s, na przypadku tablicy %s,  wielkości %d linii", sortingMethodChosenName, pivotChosenName, typeName, caseTypeChosen, ifh.Lines)
			ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("%s%sPiv%s%s-%d.csv", sortingMethodChosenName, caseTypeChosen, pivotChosenName, typeName, i)}
			ofh.InitializeSlices(repsChosen)
			list := ifh.GetListOfType(typeChosen)
			lines, times := runner.Run(sortingMethodChosen, typeChosen, ssGapChosen, pivotChosen, repsChosen, list)
			ofh.AddResultLists(lines, times)
			ofh.WriteToFile()
		}
	} else if sortingMethodChosen == 3 {
		for i := 0; i < 7; i++ {

			ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("%s%s-%d.txt", caseTypeChosen, typeName, i)}
			ifh.ReadFile()
			log.Printf("Sortowanie %s z przerwą %s dla typu %s, na przypadku tablicy %s,  wielkości %d linii", sortingMethodChosenName, ssGapChosenName, typeName, caseTypeChosen, ifh.Lines)
			ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("%s%s%s%s-%d.csv", sortingMethodChosenName, ssGapChosenName, caseTypeChosen, typeName, i)}
			ofh.InitializeSlices(repsChosen)
			list := ifh.GetListOfType(typeChosen)
			lines, times := runner.Run(sortingMethodChosen, typeChosen, ssGapChosen, pivotChosen, repsChosen, list)
			ofh.AddResultLists(lines, times)
			ofh.WriteToFile()

		}
	} else {
		for i := 0; i < 7; i++ {

			ifh := fileHandler.InputFileHandler{FileName: fmt.Sprintf("%s%s-%d.txt", caseTypeChosen, typeName, i)}
			ifh.ReadFile()
			log.Printf("Sortowanie %s dla typu %s, na przypadku tablicy %s, wielkości %d linii", sortingMethodChosenName, typeName, caseTypeChosen, ifh.Lines)
			ofh := fileHandler.OutputFileHandler{FileName: fmt.Sprintf("%s%s%s-%d.csv", sortingMethodChosenName, caseTypeChosen, typeName, i)}
			ofh.InitializeSlices(repsChosen)
			list := ifh.GetListOfType(typeChosen)
			lines, times := runner.Run(sortingMethodChosen, typeChosen, ssGapChosen, pivotChosen, repsChosen, list)
			ofh.AddResultLists(lines, times)
			ofh.WriteToFile()

		}
	}

	log.Printf("Benchmark dla %s zakończony", sortingMethodChosenName)
}
