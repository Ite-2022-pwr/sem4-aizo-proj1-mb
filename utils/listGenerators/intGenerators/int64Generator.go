package intGenerators

import (
	"math/rand"
	"sort"
)

func GenerateRandomInt64List(length int) (output []int64) {
	output = make([]int64, length)
	for i := 0; i < length; i++ {
		output[i] = int64(rand.Int() % 1000000)
	}
	return output
}

func GenerateSortedInt64List(length int) (output []int64) {
	output = make([]int64, length)
	for i, _ := range output {
		output[i] = int64(i)
	}
	return output
}

func GenerateReverseSortedInt64List(length int) (output []int64) {
	output = make([]int64, length)
	for i, _ := range output {
		output[i] = int64(length - i)
	}
	return output
}

func GenerateThirdSortedInt64List(length int) (output []int64) {
	output = GenerateRandomInt64List(length)
	sortEnd := length / 3
	sort.Slice(output[:sortEnd], func(i, j int) bool { return output[i] < output[j] })
	return output
}

func GenerateTwoThirdsSortedInt64List(length int) (output []int64) {
	output = GenerateRandomInt64List(length)
	sortEnd := 2 * (length / 3)
	sort.Slice(output[:sortEnd], func(i, j int) bool { return output[i] < output[j] })
	return output
}
