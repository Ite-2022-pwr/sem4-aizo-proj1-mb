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

func SortingMethodNameChooser(sortingMethodChosen string) (sortingMethodChosenName string) {
	switch sortingMethodChosen {
	case "q":
		sortingMethodChosenName = "quicksort"
	case "h":
		sortingMethodChosenName = "heapsort"
	case "i":
		sortingMethodChosenName = "insertionsort"
	case "s":
		sortingMethodChosenName = "shellsort"
	default:
		sortingMethodChosenName = "quicksort"
	}
	return sortingMethodChosenName
}
