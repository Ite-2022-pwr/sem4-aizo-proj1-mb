package fileHandler

import (
	"bufio"
	"log"
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
	scanner.Scan()
	fh.Lines, _ = strconv.Atoi(scanner.Text())
	log.Printf("%d lines in file", fh.Lines)
	for i := 0; i < fh.Lines; i++ {
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
