package fileHandler

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strconv"
)

type InputFileHandler struct {
	FileName string
	file     *os.File
	List     []string
	Lines    int
}

func (fh *InputFileHandler) ReadFile() {
	fh.file, _ = os.Open(fh.FileName)
	scanner := bufio.NewScanner(fh.file)
	firstLineRead := false
	for scanner.Scan() {
		if !firstLineRead {
			fh.Lines, _ = strconv.Atoi(scanner.Text())
			log.Printf("%d lines in file", fh.Lines)
			firstLineRead = true
			continue

		}
		lineStr := scanner.Text()
		fh.List = append(fh.List, lineStr)
	}
}

func (fh *InputFileHandler) GetIntList() []int {
	var intList []int
	for i := 0; i < len(fh.List); i++ {
		line := fh.List[i]
		num, _ := strconv.Atoi(line)
		intList = append(intList, num)
	}
	return intList
}

func (fh *InputFileHandler) GetFloat64List() []float64 {
	var floatList []float64
	for i := 0; i < len(fh.List); i++ {
		line := fh.List[i]
		num, _ := strconv.ParseFloat(line, 64)
		floatList = append(floatList, num)
	}
	return floatList
}

func (fh *InputFileHandler) GetFloat32List() []float32 {
	var floatList []float32
	for i := 0; i < len(fh.List); i++ {
		line := fh.List[i]
		num, _ := strconv.ParseFloat(line, 32)
		floatList = append(floatList, float32(num))
	}
	return floatList
}

func (fh *InputFileHandler) GetInt64List() []int64 {
	var intList []int64
	for i := 0; i < len(fh.List); i++ {
		line := fh.List[i]
		num, _ := strconv.ParseInt(line, 10, 64)
		intList = append(intList, num)
	}
	return intList
}

func (fh *InputFileHandler) GetInt32List() []int32 {
	var intList []int32
	for i := 0; i < len(fh.List); i++ {
		line := fh.List[i]
		num, _ := strconv.ParseInt(line, 10, 32)
		intList = append(intList, int32(num))
	}
	return intList
}

func (fh *InputFileHandler) GetStringList() []string {
	var stringList []string
	for i := 0; i < len(fh.List); i++ {
		line := fh.List[i]
		stringList = append(stringList, line)
	}
	return stringList
}

func (fh *InputFileHandler) GetShuffledFloat64List() []float64 {
	shuffledList := make([]float64, fh.Lines)
	copy(shuffledList, fh.GetFloat64List())
	rand.Shuffle(len(shuffledList), func(i, j int) { shuffledList[i], shuffledList[j] = shuffledList[j], shuffledList[i] })
	return shuffledList
}

func (fh *InputFileHandler) GetShuffledIntList() []int {
	shuffledList := make([]int, fh.Lines)
	copy(shuffledList, fh.GetIntList())
	rand.Shuffle(len(shuffledList), func(i, j int) { shuffledList[i], shuffledList[j] = shuffledList[j], shuffledList[i] })
	return shuffledList
}

func (fh *InputFileHandler) GetShuffledInt32List() []int32 {
	shuffledList := make([]int32, fh.Lines)
	copy(shuffledList, fh.GetInt32List())
	rand.Shuffle(len(shuffledList), func(i, j int) { shuffledList[i], shuffledList[j] = shuffledList[j], shuffledList[i] })
	return shuffledList
}

func (fh *InputFileHandler) GetShuffledInt64List() []int64 {
	shuffledList := make([]int64, fh.Lines)
	copy(shuffledList, fh.GetInt64List())
	rand.Shuffle(len(shuffledList), func(i, j int) { shuffledList[i], shuffledList[j] = shuffledList[j], shuffledList[i] })
	return shuffledList
}

func (fh *InputFileHandler) GetShuffledStringList() []string {
	shuffledList := make([]string, fh.Lines)
	copy(shuffledList, fh.GetStringList())
	rand.Shuffle(len(shuffledList), func(i, j int) { shuffledList[i], shuffledList[j] = shuffledList[j], shuffledList[i] })
	return shuffledList
}

func (fh *InputFileHandler) GetShuffledFloat32List() []float32 {
	shuffledList := make([]float32, fh.Lines)
	copy(shuffledList, fh.GetFloat32List())
	rand.Shuffle(len(shuffledList), func(i, j int) { shuffledList[i], shuffledList[j] = shuffledList[j], shuffledList[i] })
	return shuffledList
}
