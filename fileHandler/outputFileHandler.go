package fileHandler

import "os"

type OutputFileHandler struct {
	FileName string
	file     *os.File
	Results  []string
}

func (fh *OutputFileHandler) addResult() {

}
