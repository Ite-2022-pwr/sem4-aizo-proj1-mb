package fileHandler

import (
	"bufio"
	"fmt"
	"os"
)

type OutputFileHandler struct {
	FileName string
	file     *os.File
	Times    []int64
	Sizes    []int
}

func (fh *OutputFileHandler) InitializeSlices(length int) {
	fh.Sizes = make([]int, length)
	fh.Times = make([]int64, length)
}

func (fh *OutputFileHandler) AddResult(timeInNs int64, listSize, rep int) {
	fh.Sizes[rep] = listSize
	fh.Times[rep] = timeInNs
}

func (fh *OutputFileHandler) WriteToFile() {
	fh.file, _ = os.Create(fh.FileName)
	writer := bufio.NewWriter(fh.file)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(fh.file)
	for i := 0; i < len(fh.Times); i++ {
		_, err := writer.WriteString(fmt.Sprintf("%d;%d\n", fh.Sizes[i], fh.Times[i]))
		if err != nil {
			return
		}
	}
	writer.Flush()
}
