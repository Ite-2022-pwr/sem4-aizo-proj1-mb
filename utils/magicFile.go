package utils

import (
	"fmt"
	"log"
	"math"
	"projekt1/utils/listGenerators"
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

func generateRandomInput() {
	for i := 0; i < 7; i++ {
		length := 10000 * int(math.Pow(2, float64(i)))
		log.Printf("Generating random input of length %d", length)
		intList := listGenerators.GenerateStdIntList(0, length)
		SaveListToFile(intList, fmt.Sprintf("randomInt-%d.txt", i))
		float64List := listGenerators.GenerateFloat64List(0, length)
		SaveListToFile(float64List, fmt.Sprintf("randomFloat64-%d.txt", i))
		float32list := listGenerators.GenerateFloat32List(0, length)
		SaveListToFile(float32list, fmt.Sprintf("randomFloat32-%d.txt", i))
		int32List := listGenerators.GenerateInt32List(0, length)
		SaveListToFile(int32List, fmt.Sprintf("randomInt32-%d.txt", i))
		int64list := listGenerators.GenerateInt64List(0, length)
		SaveListToFile(int64list, fmt.Sprintf("randomInt64-%d.txt", i))
		stringList := listGenerators.GenerateStringList(0, length)
		SaveListToFile(stringList, fmt.Sprintf("randomString-%d.txt", i))
	}
}

func generateSortedInput() {
	for i := 0; i < 7; i++ {
		length := 10000 * int(math.Pow(2, float64(i)))
		log.Printf("Generating sorted input of length %d", length)
		intList := listGenerators.GenerateStdIntList(1, length)
		SaveListToFile(intList, fmt.Sprintf("sortedInt-%d.txt", i))
		float64List := listGenerators.GenerateFloat64List(1, length)
		SaveListToFile(float64List, fmt.Sprintf("sortedFloat64-%d.txt", i))
		float32list := listGenerators.GenerateFloat32List(1, length)
		SaveListToFile(float32list, fmt.Sprintf("sortedFloat32-%d.txt", i))
		int32List := listGenerators.GenerateInt32List(1, length)
		SaveListToFile(int32List, fmt.Sprintf("sortedInt32-%d.txt", i))
		int64list := listGenerators.GenerateInt64List(1, length)
		SaveListToFile(int64list, fmt.Sprintf("sortedInt64-%d.txt", i))
		stringList := listGenerators.GenerateStringList(1, length)
		SaveListToFile(stringList, fmt.Sprintf("sortedString-%d.txt", i))
	}
}

func generateReverseSortedInput() {
	for i := 0; i < 7; i++ {
		length := 10000 * int(math.Pow(2, float64(i)))
		log.Printf("Generating reverse sorted input of length %d", length)
		intList := listGenerators.GenerateStdIntList(2, length)
		SaveListToFile(intList, fmt.Sprintf("reverseSortedInt-%d.txt", i))
		float64List := listGenerators.GenerateFloat64List(2, length)
		SaveListToFile(float64List, fmt.Sprintf("reverseSortedFloat64-%d.txt", i))
		float32list := listGenerators.GenerateFloat32List(2, length)
		SaveListToFile(float32list, fmt.Sprintf("reverseSortedFloat32-%d.txt", i))
		int32List := listGenerators.GenerateInt32List(2, length)
		SaveListToFile(int32List, fmt.Sprintf("reverseSortedInt32-%d.txt", i))
		int64list := listGenerators.GenerateInt64List(2, length)
		SaveListToFile(int64list, fmt.Sprintf("reverseSortedInt64-%d.txt", i))
		stringList := listGenerators.GenerateStringList(2, length)
		SaveListToFile(stringList, fmt.Sprintf("reverseSortedString-%d.txt", i))
	}
}

func generateThirdSortedInput() {
	for i := 0; i < 7; i++ {
		length := 10000 * int(math.Pow(2, float64(i)))
		log.Printf("Generating third sorted input of length %d", length)
		intList := listGenerators.GenerateStdIntList(3, length)
		SaveListToFile(intList, fmt.Sprintf("thirdSortedInt-%d.txt", i))
		float64List := listGenerators.GenerateFloat64List(3, length)
		SaveListToFile(float64List, fmt.Sprintf("thirdSortedFloat64-%d.txt", i))
		float32list := listGenerators.GenerateFloat32List(3, length)
		SaveListToFile(float32list, fmt.Sprintf("thirdSortedFloat32-%d.txt", i))
		int32List := listGenerators.GenerateInt32List(3, length)
		SaveListToFile(int32List, fmt.Sprintf("thirdSortedInt32-%d.txt", i))
		int64list := listGenerators.GenerateInt64List(3, length)
		SaveListToFile(int64list, fmt.Sprintf("thirdSortedInt64-%d.txt", i))
		stringList := listGenerators.GenerateStringList(3, length)
		SaveListToFile(stringList, fmt.Sprintf("thirdSortedString-%d.txt", i))
	}
}

func generateTwoThirdsSortedInput() {
	for i := 0; i < 7; i++ {
		length := 10000 * int(math.Pow(2, float64(i)))
		log.Printf("Generating two thirds sorted input of length %d", length)
		intList := listGenerators.GenerateStdIntList(4, length)
		SaveListToFile(intList, fmt.Sprintf("twoThirdsSortedInt-%d.txt", i))
		float64List := listGenerators.GenerateFloat64List(4, length)
		SaveListToFile(float64List, fmt.Sprintf("twoThirdsSortedFloat64-%d.txt", i))
		float32list := listGenerators.GenerateFloat32List(4, length)
		SaveListToFile(float32list, fmt.Sprintf("twoThirdsSortedFloat32-%d.txt", i))
		int32List := listGenerators.GenerateInt32List(4, length)
		SaveListToFile(int32List, fmt.Sprintf("twoThirdsSortedInt32-%d.txt", i))
		int64list := listGenerators.GenerateInt64List(4, length)
		SaveListToFile(int64list, fmt.Sprintf("twoThirdsSortedInt64-%d.txt", i))
		stringList := listGenerators.GenerateStringList(4, length)
		SaveListToFile(stringList, fmt.Sprintf("twoThirdsSortedString-%d.txt", i))
	}
}
