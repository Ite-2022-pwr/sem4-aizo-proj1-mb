package utils

import (
	"fmt"
	"log"
	"sort"
)

func GenerateInputsForTests() {
	generateSortedInput()
	log.Print("Sorted input generated")
	generateReverseSortedInput()
	log.Print("Reverse sorted input generated")
	generateRandomInput()
	log.Print("Random input generated")
	generateThirdSortedInput()
	log.Print("Third sorted input generated")
	generateTwoThirdsSortedInput()
	log.Print("Two thirds sorted input generated")
}

func generateSortedInput() {
	for i := 1; i <= 7; i++ {
		length := 10000 * i
		log.Printf("Generating sorted input of length %d", length)
		intList := make([]int, length)
		for j := 0; j < length; j++ {
			intList[j] = j
		}
		SaveListToFile(intList, fmt.Sprintf("sortedInt-%d.txt", i))
		float64List := make([]float64, length)
		float64List = GenerateRandomFloat64List(length)
		sort.Float64s(float64List)
		SaveListToFile(float64List, fmt.Sprintf("sortedFloat64-%d.txt", i))
		float32list := make([]float32, length)
		float32list = GenerateRandomFloat32List(length)
		sort.Slice(float32list, func(i, j int) bool { return float32list[i] < float32list[j] })
		SaveListToFile(float32list, fmt.Sprintf("sortedFloat32-%d.txt", i))
		int32List := make([]int32, length)
		for j := 0; j < length; j++ {
			int32List[j] = int32(j)
		}
		SaveListToFile(int32List, fmt.Sprintf("sortedInt32-%d.txt", i))
		int64list := make([]int64, length)
		for j := 0; j < length; j++ {
			int64list[j] = int64(j)
		}
		SaveListToFile(int64list, fmt.Sprintf("sortedInt64-%d.txt", i))
		stringList := make([]string, length)
		stringList = GenerateRandomStringList(length)
		sort.Strings(stringList)
		SaveListToFile(stringList, fmt.Sprintf("sortedString-%d.txt", i))
	}
}

func generateReverseSortedInput() {
	for i := 1; i <= 7; i++ {
		length := 10000 * i
		log.Printf("Generating reverse sorted input of length %d", length)
		intList := make([]int, length)
		for j := 0; j < length; j++ {
			intList[j] = length - j
		}
		SaveListToFile(intList, fmt.Sprintf("reverseSortedInt-%d.txt", i))
		float64List := make([]float64, length)
		float64List = GenerateRandomFloat64List(length)
		sort.Sort(sort.Reverse(sort.Float64Slice(float64List)))
		SaveListToFile(float64List, fmt.Sprintf("reverseSortedFloat64-%d.txt", i))
		float32list := make([]float32, length)
		float32list = GenerateRandomFloat32List(length)
		sort.Slice(float32list, func(i, j int) bool { return float32list[i] > float32list[j] })
		SaveListToFile(float32list, fmt.Sprintf("reverseSortedFloat32-%d.txt", i))
		int32List := make([]int32, length)
		for j := 0; j < length; j++ {
			int32List[j] = int32(length - j)
		}
		SaveListToFile(int32List, fmt.Sprintf("reverseSortedInt32-%d.txt", i))
		int64list := make([]int64, length)
		for j := 0; j < length; j++ {
			int64list[j] = int64(length - j)
		}
		SaveListToFile(int64list, fmt.Sprintf("reverseSortedInt64-%d.txt", i))
		stringList := make([]string, length)
		stringList = GenerateRandomStringList(length)
		sort.Strings(stringList)
		last := len(stringList) - 1
		half := len(stringList) / 2
		for j := 0; j < half; j++ {
			stringList[j], stringList[last-j] = stringList[last-j], stringList[j]
		}
		SaveListToFile(stringList, fmt.Sprintf("reverseSortedString-%d.txt", i))
	}
}

func generateRandomInput() {
	for i := 1; i <= 7; i++ {
		length := 10000 * i
		log.Printf("Generating random input of length %d", length)
		intList := make([]int, length)
		intList = GenerateRandomIntList(length)
		SaveListToFile(intList, fmt.Sprintf("randomInt-%d.txt", i))
		float64List := make([]float64, length)
		float64List = GenerateRandomFloat64List(length)
		SaveListToFile(float64List, fmt.Sprintf("randomFloat64-%d.txt", i))
		float32list := make([]float32, length)
		float32list = GenerateRandomFloat32List(length)
		SaveListToFile(float32list, fmt.Sprintf("randomFloat32-%d.txt", i))
		int32List := make([]int32, length)
		int32List = GenerateRandomInt32List(length)
		SaveListToFile(int32List, fmt.Sprintf("randomInt32-%d.txt", i))
		int64list := make([]int64, length)
		int64list = GenerateRandomInt64List(length)
		SaveListToFile(int64list, fmt.Sprintf("randomInt64-%d.txt", i))
		stringList := make([]string, length)
		stringList = GenerateRandomStringList(length)
		SaveListToFile(stringList, fmt.Sprintf("randomString-%d.txt", i))
	}
}

func generateThirdSortedInput() {
	for i := 1; i <= 7; i++ {
		length := 10000 * i
		log.Printf("Generating third sorted input of length %d", length)
		sortEnd := length / 3
		intList := make([]int, length)
		intList = GenerateRandomIntList(length)
		sort.Slice(intList[0:sortEnd], func(i, j int) bool { return intList[i] < intList[j] })
		SaveListToFile(intList, fmt.Sprintf("thirdSortedInt-%d.txt", i))
		float64List := make([]float64, length)
		float64List = GenerateRandomFloat64List(length)
		sort.Slice(float64List[0:sortEnd], func(i, j int) bool { return float64List[i] < float64List[j] })
		SaveListToFile(float64List, fmt.Sprintf("thirdSortedFloat64-%d.txt", i))
		float32list := make([]float32, length)
		float32list = GenerateRandomFloat32List(length)
		sort.Slice(float32list[0:sortEnd], func(i, j int) bool { return float32list[i] < float32list[j] })
		SaveListToFile(float32list, fmt.Sprintf("thirdSortedFloat32-%d.txt", i))
		int32List := make([]int32, length)
		int32List = GenerateRandomInt32List(length)
		sort.Slice(int32List[0:sortEnd], func(i, j int) bool { return int32List[i] < int32List[j] })
		SaveListToFile(int32List, fmt.Sprintf("thirdSortedInt32-%d.txt", i))
		int64list := make([]int64, length)
		int64list = GenerateRandomInt64List(length)
		sort.Slice(int64list[0:sortEnd], func(i, j int) bool { return int64list[i] < int64list[j] })
		SaveListToFile(int64list, fmt.Sprintf("thirdSortedInt64-%d.txt", i))
		stringList := make([]string, length)
		stringList = GenerateRandomStringList(length)
		sort.Strings(stringList[0:sortEnd])
		SaveListToFile(stringList, fmt.Sprintf("thirdSortedString-%d.txt", i))
	}
}

func generateTwoThirdsSortedInput() {
	for i := 1; i <= 7; i++ {
		length := 10000 * i
		log.Printf("Generating two thirds sorted input of length %d", length)
		sortEnd := (length / 3) * 2
		intList := make([]int, length)
		intList = GenerateRandomIntList(length)
		sort.Slice(intList[0:sortEnd], func(i, j int) bool { return intList[i] < intList[j] })
		SaveListToFile(intList, fmt.Sprintf("twoThirdsSortedInt-%d.txt", i))
		float64List := make([]float64, length)
		float64List = GenerateRandomFloat64List(length)
		sort.Slice(float64List[0:sortEnd], func(i, j int) bool { return float64List[i] < float64List[j] })
		SaveListToFile(float64List, fmt.Sprintf("twoThirdsSortedFloat64-%d.txt", i))
		float32list := make([]float32, length)
		float32list = GenerateRandomFloat32List(length)
		sort.Slice(float32list[0:sortEnd], func(i, j int) bool { return float32list[i] < float32list[j] })
		SaveListToFile(float32list, fmt.Sprintf("twoThirdsSortedFloat32-%d.txt", i))
		int32List := make([]int32, length)
		int32List = GenerateRandomInt32List(length)
		sort.Slice(int32List[0:sortEnd], func(i, j int) bool { return int32List[i] < int32List[j] })
		SaveListToFile(int32List, fmt.Sprintf("twoThirdsSortedInt32-%d.txt", i))
		int64list := make([]int64, length)
		int64list = GenerateRandomInt64List(length)
		sort.Slice(int64list[0:sortEnd], func(i, j int) bool { return int64list[i] < int64list[j] })
		SaveListToFile(int64list, fmt.Sprintf("twoThirdsSortedInt64-%d.txt", i))
		stringList := make([]string, length)
		stringList = GenerateRandomStringList(length)
		sort.Strings(stringList[0:sortEnd])
		SaveListToFile(stringList, fmt.Sprintf("twoThirdsSortedString-%d.txt", i))
	}
}
