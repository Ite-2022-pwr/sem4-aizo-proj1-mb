package listGenerators

import "projekt1/utils/listGenerators/intGenerators"

func GenerateStdIntList(generationStrat, length int) (output []int) {
	switch generationStrat {
	case 0:
		output = intGenerators.GenerateRandomIntList(length)
	case 1:
		output = intGenerators.GenerateSortedIntList(length)
	case 2:
		output = intGenerators.GenerateReverseSortedIntList(length)
	case 3:
		output = intGenerators.GenerateThirdSortedIntList(length)
	case 4:
		output = intGenerators.GenerateTwoThirdsSortedIntList(length)
	default:
		output = intGenerators.GenerateRandomIntList(length)
	}
	return output
}

func GenerateInt32List(generationStrat, length int) (output []int32) {
	switch generationStrat {
	case 0:
		output = intGenerators.GenerateRandomInt32List(length)
	case 1:
		output = intGenerators.GenerateSortedInt32List(length)
	case 2:
		output = intGenerators.GenerateReverseSortedInt32List(length)
	case 3:
		output = intGenerators.GenerateThirdSortedInt32List(length)
	case 4:
		output = intGenerators.GenerateTwoThirdsSortedInt32List(length)
	default:
		output = intGenerators.GenerateRandomInt32List(length)
	}
	return output
}

func GenerateInt64List(generationStrat, length int) (output []int64) {
	switch generationStrat {
	case 0:
		output = intGenerators.GenerateRandomInt64List(length)
	case 1:
		output = intGenerators.GenerateSortedInt64List(length)
	case 2:
		output = intGenerators.GenerateReverseSortedInt64List(length)
	case 3:
		output = intGenerators.GenerateThirdSortedInt64List(length)
	case 4:
		output = intGenerators.GenerateTwoThirdsSortedInt64List(length)
	default:
		output = intGenerators.GenerateRandomInt64List(length)
	}
	return output
}
