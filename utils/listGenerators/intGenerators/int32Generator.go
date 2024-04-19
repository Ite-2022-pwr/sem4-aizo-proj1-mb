package intGenerators

import (
	"math/rand"
	"sort"
)

func GenerateRandomInt32List(length int) []int32 {
	output := make([]int32, length)
	for i := 0; i < length; i++ {
		output[i] = int32(rand.Int() % 1000000)
	}
	return output
}

func GenerateSortedInt32List(length int) []int32 {
	output := make([]int32, length)
	for i := range output {
		output[i] = int32(i)
	}
	return output
}

func GenerateReverseSortedInt32List(length int) []int32 {
	output := make([]int32, length)
	for i := range output {
		output[i] = int32(length - i)
	}
	return output
}

func GenerateThirdSortedInt32List(length int) []int32 {
	output := GenerateRandomInt32List(length)
	sortEnd := length / 3
	sort.Slice(output[:sortEnd], func(i, j int) bool { return output[i] < output[j] })
	return output
}

func GenerateTwoThirdsSortedInt32List(length int) []int32 {
	output := GenerateRandomInt32List(length)
	sortEnd := 2 * (length / 3)
	sort.Slice(output[:sortEnd], func(i, j int) bool { return output[i] < output[j] })
	return output
}
