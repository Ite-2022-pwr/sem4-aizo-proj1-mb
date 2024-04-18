package floatGenerators

import (
	"math/rand"
	"sort"
)

func GenerateRandomFloat64List(length int) (output []float64) {
	output = make([]float64, length)
	for i, _ := range output {
		output[i] = rand.Float64()
	}
	return output
}

func GenerateSortedFloat64List(length int) (output []float64) {
	output = make([]float64, length)
	for i, _ := range output {
		output[i] = float64(i) / 10000
	}
	return output
}

func GenerateReverseSortedFloat64List(length int) (output []float64) {
	output = make([]float64, length)
	for i, _ := range output {
		output[i] = float64(length-i) / 10000
	}
	return output
}

func GenerateThirdSortedFloat64List(length int) (output []float64) {
	output = GenerateRandomFloat64List(length)
	sortEnd := length / 3
	sort.Slice(output[:sortEnd], func(i, j int) bool { return output[i] < output[j] })
	return output
}

func GenerateTwoThirdsSortedFloat64List(length int) (output []float64) {
	output = GenerateRandomFloat64List(length)
	sortEnd := 2 * (length / 3)
	sort.Slice(output[:sortEnd], func(i, j int) bool { return output[i] < output[j] })
	return output
}
