package submenus

import (
	"bufio"
	"fmt"
	"os"
	"projekt1/fileHandler"
	"projekt1/utils"
	"strings"
)

func OutputListToFile(list []any) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Podaj nazwę pliku do zapisania listy:")
	outputFilename, _ := reader.ReadString('\n')
	outputFilename = strings.TrimSuffix(outputFilename, "\n")
	utils.SaveListToFile(list, outputFilename)
}

func SaveResultToFile(lines []int, times []int64) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Podaj nazwę pliku do zapisania wyniku:")
	outputFilename, _ := reader.ReadString('\n')
	outputFilename = strings.TrimSuffix(outputFilename, "\n")
	ofh := fileHandler.OutputFileHandler{FileName: outputFilename}
	ofh.AddResultLists(lines, times)
	ofh.WriteToFile()
}
