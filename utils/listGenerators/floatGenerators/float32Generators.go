package floatGenerators

import (
	"math/rand"
	"sort"
)

func GenerateRandomFloat32List(length int) (output []float32) {
	output = make([]float32, length)
	for i := range output {
		output[i] = rand.Float32()
	}
	return output
}

func GenerateSortedFloat32List(length int) (output []float32) {
	output = make([]float32, length)
	for i := range output {
		output[i] = float32(i) / 1000
	}
	return output
}

func GenerateReverseSortedFloat32List(length int) (output []float32) {
	output = make([]float32, length)
	for i := range output {
		output[i] = float32(length-i) / 1000
	}
	return output
}

func GenerateThirdSortedFloat32List(length int) (output []float32) {
	output = GenerateRandomFloat32List(length)
	sortEnd := length / 3
	sort.Slice(output[:sortEnd], func(i, j int) bool { return output[i] < output[j] })
	return output
}

func GenerateTwoThirdsSortedFloat32List(length int) (output []float32) {
	output = GenerateRandomFloat32List(length)
	sortEnd := 2 * (length / 3)
	sort.Slice(output[:sortEnd], func(i, j int) bool { return output[i] < output[j] })
	return output
}
