package submenus

import (
	"bufio"
	"fmt"
	"os"
	"projekt1/fileHandler"
	"projekt1/utils"
	"strings"
)

func InputDataFromFile() (data []any, typeChosen int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Podaj nazwę pliku do wczytania:")
	inputSourceStr, _ := reader.ReadString('\n')
	inputSourceStr = strings.TrimSuffix(inputSourceStr, "\n")
	ifh := fileHandler.InputFileHandler{FileName: inputSourceStr}
	ifh.ReadFile()
	typeChosen = ChooseType()
	switch typeChosen {
	case 0:
		data = utils.CastToAnySlice(ifh.GetIntList())
	case 1:
		data = utils.CastToAnySlice(ifh.GetFloat64List())
	case 2:
		data = utils.CastToAnySlice(ifh.GetFloat32List())
	case 3:
		data = utils.CastToAnySlice(ifh.GetInt32List())
	case 4:
		data = utils.CastToAnySlice(ifh.GetInt64List())
	case 5:
		data = utils.CastToAnySlice(ifh.GetStringList())
	default:
		data = utils.CastToAnySlice(ifh.GetIntList())
	}
	return data, typeChosen
}
