package utils

func PivotNameChooser(pivotChosen int) (pivotChosenName string) {
	switch pivotChosen {
	case 0:
		pivotChosenName = "High"
	case 1:
		pivotChosenName = "Low"
	case 2:
		pivotChosenName = "Med"
	case 3:
		pivotChosenName = "Rand"
	default:
		pivotChosenName = "High"
	}
	return pivotChosenName
}

func TypeNameChooser(typeChosen int) (typeName string) {

	switch typeChosen {
	case 0:
		typeName = "Int"
	case 1:
		typeName = "Float64"
	case 2:
		typeName = "Float32"
	case 3:
		typeName = "Int32"
	case 4:
		typeName = "Int64"
	case 5:
		typeName = "String"
	default:
		typeName = "Int"
	}
	return typeName
}

func SsGapNameChooser(ssGapChosen int) (ssGapName string) {
	switch ssGapChosen {
	case 0:
		ssGapName = "shell"
	case 1:
		ssGapName = "frankLazarus"
	default:
		ssGapName = "shell"
	}
	return ssGapName
}

func SortingMethodNameChooser(sortingMethodChosen int) (sortingMethodChosenName string) {
	switch sortingMethodChosen {
	case 0:
		sortingMethodChosenName = "quicksort"
	case 1:
		sortingMethodChosenName = "heapsort"
	case 2:
		sortingMethodChosenName = "insertionsort"
	case 3:
		sortingMethodChosenName = "shellsort"
	default:
		sortingMethodChosenName = "quicksort"
	}
	return sortingMethodChosenName
}

func CaseTypeNameChooser(caseTypeChosen int) (caseTypeChosenName string) {
	switch caseTypeChosen {
	case 0:
		caseTypeChosenName = "random"
	case 1:
		caseTypeChosenName = "sorted"
	case 2:
		caseTypeChosenName = "reverseSorted"
	case 3:
		caseTypeChosenName = "thirdSorted"
	case 4:
		caseTypeChosenName = "twoThirdsSorted"
	default:
		caseTypeChosenName = "random"
	}
	return caseTypeChosenName
}
