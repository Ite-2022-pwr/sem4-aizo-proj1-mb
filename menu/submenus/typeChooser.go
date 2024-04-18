package submenus

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChooseType() (typeChosen int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Wybierz typ zmiennej:")
	fmt.Println("0 - int")
	fmt.Println("1 - float64")
	fmt.Println("2 - float32")
	fmt.Println("3 - int32")
	fmt.Println("4 - int64")
	fmt.Println("5 - string")
	typeChosenStr, _ := reader.ReadString('\n')
	typeChosenStr = strings.TrimSuffix(typeChosenStr, "\n")
	typeChosen, _ = strconv.Atoi(typeChosenStr)
	return typeChosen
}
