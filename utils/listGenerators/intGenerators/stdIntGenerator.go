package intGenerators

import (
	"math/rand"
	"sort"
)

func GenerateRandomIntList(length int) (output []int) {
	output = make([]int, length)
	for i := 0; i < length; i++ {
		output[i] = rand.Int() % 100000
	}
	return output
}

func GenerateSortedIntList(length int) (output []int) {
	output = make([]int, length)
	for i := range output {
		output[i] = i
	}
	return output
}

func GenerateReverseSortedIntList(length int) (output []int) {
	output = make([]int, length)
	for i := range output {
		output[i] = length - i
	}
	return output
}

func GenerateThirdSortedIntList(length int) (output []int) {
	output = GenerateRandomIntList(length)
	sortEnd := length / 3
	sort.Ints(output[:sortEnd])
	return output
}

func GenerateTwoThirdsSortedIntList(length int) (output []int) {
	output = GenerateRandomIntList(length)
	sortEnd := 2 * (length / 3)
	sort.Ints(output[:sortEnd])
	return output
}
