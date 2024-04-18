package submenus

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChooseGenerationStrategy() (generationStrat int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Wybierz sposób generacji danych:")
	fmt.Println("0 - losowe")
	fmt.Println("1 - posortowane rosnąco")
	fmt.Println("2 - posortowane malejąco")
	fmt.Println("3 - posortowane rosnąco w jednej trzeciej")
	fmt.Println("4 - posortowane rosnąco w dwóch trzecich")
	generationStratStr, _ := reader.ReadString('\n')
	generationStratStr = strings.TrimSuffix(generationStratStr, "\n")
	generationStrat, _ = strconv.Atoi(generationStratStr)
	return generationStrat
}
