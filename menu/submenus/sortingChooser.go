package submenus

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChooseSorting() (sortingChosen, ssGapChosen, qsPivotChosen int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Wybierz sposób sortowania:")
	fmt.Println("0 - quicksort")
	fmt.Println("1 - heapsort")
	fmt.Println("2 - insertionsort")
	fmt.Println("3 - shellsort")
	sortingChosenStr, _ := reader.ReadString('\n')
	sortingChosen, _ = strconv.Atoi(strings.TrimSuffix(sortingChosenStr, "\n"))
	if sortingChosen == 3 {
		fmt.Println("Wybierz sposób wyboru odstępów:")
		fmt.Println("0 - shell")
		fmt.Println("1 - frankLazarus")
		ssGapChosenStr, _ := reader.ReadString('\n')
		ssGapChosen, _ = strconv.Atoi(strings.TrimSuffix(ssGapChosenStr, "\n"))
	}
	if sortingChosen == 0 {
		fmt.Println("Wybierz sposób wyboru pivota:")
		fmt.Println("0 - High")
		fmt.Println("1 - Low")
		fmt.Println("2 - Med")
		fmt.Println("3 - Rand")
		qsPivotChosenStr, _ := reader.ReadString('\n')
		qsPivotChosen, _ = strconv.Atoi(strings.TrimSuffix(qsPivotChosenStr, "\n"))
	}
	return sortingChosen, ssGapChosen, qsPivotChosen
}
