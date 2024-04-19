package stringGenerators

import (
	"math/rand"
	"sort"
)

func GenerateRandomStringList(length int) (output []string) {
	output = make([]string, length)
	charRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for i := range output {
		b := make([]rune, (rand.Int()%16)+1)
		for i := range b {
			b[i] = charRunes[rand.Intn(len(charRunes))]
		}
		tempStr := string(b)
		output[i] = tempStr
	}
	return output
}

func GenerateSortedStringList(length int) (output []string) {
	output = GenerateRandomStringList(length)
	sort.Strings(output)
	return output
}

func GenerateReverseSortedStringList(length int) (output []string) {
	output = GenerateRandomStringList(length)
	sort.Sort(sort.Reverse(sort.StringSlice(output)))
	return output
}

func GenerateThirdSortedStringList(length int) (output []string) {
	output = GenerateRandomStringList(length)
	sortEnd := length / 3
	sort.Strings(output[:sortEnd])
	return output
}

func GenerateTwoThirdsSortedStringList(length int) (output []string) {
	output = GenerateRandomStringList(length)
	sortEnd := 2 * (length / 3)
	sort.Strings(output[:sortEnd])
	return output
}
