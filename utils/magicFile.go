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
	for i := 1; i <= 10; i++ {
		intList := make([]int, 10000*i)
		for j := 0; j < 10000*i; j++ {
			intList[j] = j
		}
		SaveListToFile(intList, fmt.Sprintf("sortedInt-%d.txt", i))
	}

	for i := 1; i <= 10; i++ {
		float64List := make([]float64, 10000*i)
		float64List = GenerateRandomFloat64List(10000 * i)
		sort.Float64s(float64List)
		SaveListToFile(float64List, fmt.Sprintf("sortedFloat64-%d.txt", i))
	}

	for i := 1; i <= 10; i++ {
		float32list := make([]float32, 10000*i)
		float32list = GenerateRandomFloat32List(10000 * i)
		sort.Slice(float32list, func(i, j int) bool { return float32list[i] < float32list[j] })
		SaveListToFile(float32list, fmt.Sprintf("sortedFloat32-%d.txt", i))
	}

	for i := 1; i <= 10; i++ {
		int32List := make([]int32, 10000*i)
		for j := 0; j < 10000*i; j++ {
			int32List[j] = int32(j)
		}
		SaveListToFile(int32List, fmt.Sprintf("sortedInt32-%d.txt", i))
	}

	for i := 1; i <= 10; i++ {
		int64list := make([]int64, 10000*i)
		for j := 0; j < 10000*i; j++ {
			int64list[j] = int64(j)
		}
		SaveListToFile(int64list, fmt.Sprintf("sortedInt64-%d.txt", i))
	}

	for i := 1; i <= 10; i++ {
		stringList := make([]string, 10000*i)
		stringList = GenerateRandomStringList(10000 * i)
		sort.Strings(stringList)
		SaveListToFile(stringList, fmt.Sprintf("sortedString-%d.txt", i))
	}
}

func generateReverseSortedInput() {
	for i := 1; i <= 10; i++ {
		intList := make([]int, 10000*i)
		for j := 0; j < 10000*i; j++ {
			intList[j] = 10000*i - j
		}
		SaveListToFile(intList, fmt.Sprintf("reverseSortedInt-%d.txt", i))
	}
	for i := 1; i <= 10; i++ {
		float64List := make([]float64, 10000*i)
		float64List = GenerateRandomFloat64List(10000 * i)
		sort.Sort(sort.Reverse(sort.Float64Slice(float64List)))
		SaveListToFile(float64List, fmt.Sprintf("reverseSortedFloat64-%d.txt", i))
	}
	for i := 1; i <= 10; i++ {
		float32list := make([]float32, 10000*i)
		float32list = GenerateRandomFloat32List(10000 * i)
		sort.Slice(float32list, func(i, j int) bool { return float32list[i] > float32list[j] })
		SaveListToFile(float32list, fmt.Sprintf("reverseSortedFloat32-%d.txt", i))
	}
	for i := 1; i <= 10; i++ {
		int32List := make([]int32, 10000*i)
		for j := 0; j < 10000*i; j++ {
			int32List[j] = int32(10000*i - j)
		}
		SaveListToFile(int32List, fmt.Sprintf("reverseSortedInt32-%d.txt", i))
	}
	for i := 1; i <= 10; i++ {
		int64list := make([]int64, 10000*i)
		for j := 0; j < 10000*i; j++ {
			int64list[j] = int64(10000*i - j)
		}
		SaveListToFile(int64list, fmt.Sprintf("reverseSortedInt64-%d.txt", i))
	}
	for i := 1; i <= 10; i++ {
		stringList := make([]string, 10000*i)
		stringList = GenerateRandomStringList(10000 * i)
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
	for i := 1; i <= 10; i++ {
		intList := make([]int, 10000*i)
		intList = GenerateRandomIntList(10000 * i)
		SaveListToFile(intList, fmt.Sprintf("randomInt-%d.txt", i))
	}
	for i := 1; i <= 10; i++ {
		float64List := make([]float64, 10000*i)
		float64List = GenerateRandomFloat64List(10000 * i)
		SaveListToFile(float64List, fmt.Sprintf("randomFloat64-%d.txt", i))
	}
	for i := 1; i <= 10; i++ {
		float32list := make([]float32, 10000*i)
		float32list = GenerateRandomFloat32List(10000 * i)
		SaveListToFile(float32list, fmt.Sprintf("randomFloat32-%d.txt", i))
	}
	for i := 1; i <= 10; i++ {
		int32List := make([]int32, 10000*i)
		int32List = GenerateRandomInt32List(10000 * i)
		SaveListToFile(int32List, fmt.Sprintf("randomInt32-%d.txt", i))
	}
	for i := 1; i <= 10; i++ {
		int64list := make([]int64, 10000*i)
		int64list = GenerateRandomInt64List(10000 * i)
		SaveListToFile(int64list, fmt.Sprintf("randomInt64-%d.txt", i))
	}
	for i := 1; i <= 10; i++ {
		stringList := make([]string, 10000*i)
		stringList = GenerateRandomStringList(10000 * i)
		SaveListToFile(stringList, fmt.Sprintf("randomString-%d.txt", i))
	}
}

func generateThirdSortedInput() {
	for i := 1; i <= 10; i++ {
		intList := make([]int, 3333*i)
		for j := 0; j < 3333*i; j++ {
			intList[j] = j
		}
		intList = append(intList, GenerateRandomIntList(6667*i)...)
		SaveListToFile(intList, fmt.Sprintf("thirdSortedInt-%d.txt", i))
	}
	for i := 1; i <= 10; i++ {
		float64List := make([]float64, 10000*i)
		float64List = GenerateRandomFloat64List(3333 * i)
		sort.Float64s(float64List)
		float64List = append(float64List, GenerateRandomFloat64List(6667*i)...)
		SaveListToFile(float64List, fmt.Sprintf("thirdSortedFloat64-%d.txt", i))
	}
	for i := 1; i <= 10; i++ {
		float32list := make([]float32, 10000*i)
		float32list = GenerateRandomFloat32List(3333 * i)
		sort.Slice(float32list, func(i, j int) bool { return float32list[i] < float32list[j] })
		float32list = append(float32list, GenerateRandomFloat32List(6667*i)...)
		SaveListToFile(float32list, fmt.Sprintf("thirdSortedFloat32-%d.txt", i))
	}
	for i := 1; i <= 10; i++ {
		int32List := make([]int32, 3333*i)
		sort.Slice(int32List, func(i, j int) bool { return int32List[i] < int32List[j] })
		int32List = append(int32List, GenerateRandomInt32List(6667*i)...)
		SaveListToFile(int32List, fmt.Sprintf("thirdSortedInt32-%d.txt", i))
	}
	for i := 1; i <= 10; i++ {
		int64list := make([]int64, 3333*i)
		sort.Slice(int64list, func(i, j int) bool { return int64list[i] < int64list[j] })
		int64list = append(int64list, GenerateRandomInt64List(6667*i)...)
		SaveListToFile(int64list, fmt.Sprintf("thirdSortedInt64-%d.txt", i))
	}
	for i := 1; i <= 10; i++ {
		stringList := make([]string, 10000*i)
		stringList = GenerateRandomStringList(3333 * i)
		sort.Strings(stringList)
		stringList = append(stringList, GenerateRandomStringList(6667*i)...)
		SaveListToFile(stringList, fmt.Sprintf("thirdSortedString-%d.txt", i))
	}
}

func generateTwoThirdsSortedInput() {
	for i := 1; i <= 10; i++ {
		intList := make([]int, 6667*i)
		for j := 0; j < 6667*i; j++ {
			intList[j] = j
		}
		intList = append(intList, GenerateRandomIntList(3333*i)...)
		SaveListToFile(intList, fmt.Sprintf("twoThirdsSortedInt-%d.txt", i))
	}
	for i := 1; i <= 10; i++ {
		float64List := make([]float64, 10000*i)
		float64List = GenerateRandomFloat64List(6667 * i)
		sort.Float64s(float64List)
		float64List = append(float64List, GenerateRandomFloat64List(3333*i)...)
		SaveListToFile(float64List, fmt.Sprintf("twoThirdsSortedFloat64-%d.txt", i))
	}
	for i := 1; i <= 10; i++ {
		float32list := make([]float32, 10000*i)
		float32list = GenerateRandomFloat32List(6667 * i)
		sort.Slice(float32list, func(i, j int) bool { return float32list[i] < float32list[j] })
		float32list = append(float32list, GenerateRandomFloat32List(3333*i)...)
		SaveListToFile(float32list, fmt.Sprintf("twoThirdsSortedFloat32-%d.txt", i))
	}
	for i := 1; i <= 10; i++ {
		int32List := make([]int32, 6667*i)
		sort.Slice(int32List, func(i, j int) bool { return int32List[i] < int32List[j] })
		int32List = append(int32List, GenerateRandomInt32List(3333*i)...)
		SaveListToFile(int32List, fmt.Sprintf("twoThirdsSortedInt32-%d.txt", i))
	}
	for i := 1; i <= 10; i++ {
		int64list := make([]int64, 6667*i)
		sort.Slice(int64list, func(i, j int) bool { return int64list[i] < int64list[j] })
		int64list = append(int64list, GenerateRandomInt64List(3333*i)...)
		SaveListToFile(int64list, fmt.Sprintf("twoThirdsSortedInt64-%d.txt", i))
	}
	for i := 1; i <= 10; i++ {
		stringList := make([]string, 10000*i)
		stringList = GenerateRandomStringList(6667 * i)
		sort.Strings(stringList)
		stringList = append(stringList, GenerateRandomStringList(3333*i)...)
		SaveListToFile(stringList, fmt.Sprintf("twoThirdsSortedString-%d.txt", i))
	}
}