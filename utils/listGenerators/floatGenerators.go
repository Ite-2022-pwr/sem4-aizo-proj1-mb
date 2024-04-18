package listGenerators

import "projekt1/utils/listGenerators/floatGenerators"

func GenerateFloat64List(generationStrat, length int) (output []float64) {
	switch generationStrat {
	case 0:
		output = floatGenerators.GenerateRandomFloat64List(length)
	case 1:
		output = floatGenerators.GenerateSortedFloat64List(length)
	case 2:
		output = floatGenerators.GenerateReverseSortedFloat64List(length)
	case 3:
		output = floatGenerators.GenerateThirdSortedFloat64List(length)
	case 4:
		output = floatGenerators.GenerateTwoThirdsSortedFloat64List(length)
	default:
		output = floatGenerators.GenerateRandomFloat64List(length)
	}
	return output
}

func GenerateFloat32List(generationStrat, length int) (output []float32) {
	switch generationStrat {
	case 0:
		output = floatGenerators.GenerateRandomFloat32List(length)
	case 1:
		output = floatGenerators.GenerateSortedFloat32List(length)
	case 2:
		output = floatGenerators.GenerateReverseSortedFloat32List(length)
	case 3:
		output = floatGenerators.GenerateThirdSortedFloat32List(length)
	case 4:
		output = floatGenerators.GenerateTwoThirdsSortedFloat32List(length)
	default:
		output = floatGenerators.GenerateRandomFloat32List(length)
	}
	return output
}
