package listGenerators

import "projekt1/utils/listGenerators/stringGenerators"

func GenerateStringList(generationStrat, length int) (output []string) {
	switch generationStrat {
	case 0:
		output = stringGenerators.GenerateRandomStringList(length)
	case 1:
		output = stringGenerators.GenerateSortedStringList(length)
	case 2:
		output = stringGenerators.GenerateReverseSortedStringList(length)
	case 3:
		output = stringGenerators.GenerateThirdSortedStringList(length)
	case 4:
		output = stringGenerators.GenerateTwoThirdsSortedStringList(length)
	default:
		output = stringGenerators.GenerateRandomStringList(length)
	}
	return output
}
