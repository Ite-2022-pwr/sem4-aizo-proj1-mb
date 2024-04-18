package submenus

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChooseReps() (reps int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Podaj liczbę powtórzeń:")
	repsStr, _ := reader.ReadString('\n')
	reps, _ = strconv.Atoi(strings.TrimSuffix(repsStr, "\n"))
	return reps
}
