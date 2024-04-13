package fileHandler

import "os"

type OutputFileHandler struct {
	FileName string
	file     *os.File
	Times    []int
	Sizes    []int
}

func (fh *OutputFileHandler) AddResult(timeInNs, listSize int) {
	fh.Sizes = append(fh.Sizes, listSize)
	fh.Times = append(fh.Times, timeInNs)
}
