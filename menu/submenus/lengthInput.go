package submenus

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func EnterLength() (length int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Podaj długość listy:")
	lengthStr, _ := reader.ReadString('\n')
	lengthStr = strings.TrimSuffix(lengthStr, "\n")
	length, _ = strconv.Atoi(lengthStr)
	return length

}
