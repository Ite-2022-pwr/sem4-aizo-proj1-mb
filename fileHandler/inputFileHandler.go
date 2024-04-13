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

func (fh *InputFileHandler) GetFloatList() []float64 {
	var floatList []float64
	for i := 0; i < len(fh.List); i++ {
		line := fh.List[i]
		num, _ := strconv.ParseFloat(line, 64)
		floatList = append(floatList, num)
	}
	return floatList
}

func (fh *InputFileHandler) GetShuffledFloatList() []float64 {
	shuffledList := make([]float64, fh.Lines)
	copy(shuffledList, fh.GetFloatList())
	rand.Shuffle(len(shuffledList), func(i, j int) { shuffledList[i], shuffledList[j] = shuffledList[j], shuffledList[i] })
	return shuffledList
}

func (fh *InputFileHandler) GetShuffledIntList() []int {
	shuffledList := make([]int, fh.Lines)
	copy(shuffledList, fh.GetIntList())
	rand.Shuffle(len(shuffledList), func(i, j int) { shuffledList[i], shuffledList[j] = shuffledList[j], shuffledList[i] })
	return shuffledList
}
