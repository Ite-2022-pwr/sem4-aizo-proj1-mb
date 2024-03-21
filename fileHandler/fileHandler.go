package fileHandler

import (
	"bufio"
	"os"
	"strconv"
)

type FileHandler struct {
	FileName string
	file     *os.File
	List     []string
}

func (fh *FileHandler) ReadFile() {
	fh.file, _ = os.Open(fh.FileName)
	scanner := bufio.NewScanner(fh.file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		fh.List = append(fh.List, lineStr)
	}
}

func (fh *FileHandler) GetIntList() []int {
	var intList []int
	for i := 0; i < len(fh.List); i++ {
		line := fh.List[i]
		num, _ := strconv.Atoi(line)
		intList = append(intList, num)
	}
	return intList
}

func (fh *FileHandler) GetFloatList() []float64 {
	var floatList []float64
	for i := 0; i < len(fh.List); i++ {
		line := fh.List[i]
		num, _ := strconv.ParseFloat(line, 64)
		floatList = append(floatList, num)
	}
	return floatList
}
