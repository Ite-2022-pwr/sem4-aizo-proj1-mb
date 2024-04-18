package submenus

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChooseAction() (actionChosen int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Wybierz akcję:")
	fmt.Println("0 - wczytaj dane z pliku")
	fmt.Println("1 - wygeneruj dane losowe")
	fmt.Println("2 - zapisz ostatnio wygenerowaną listę do pliku")
	fmt.Println("3 - sortuj listę")
	fmt.Println("4 - wyświetl ostatnio wygenerowaną listę")
	fmt.Println("5 - wybierz algorytmy")
	fmt.Println("6 - wybierz liczbę powtórzeń")
	fmt.Println("7 - zapisz ostatni wynik do pliku")
	fmt.Println("8 - wyświetl wybrane parametry")
	fmt.Println("9 - wyjście")
	actionChosenStr, _ := reader.ReadString('\n')
	actionChosen, _ = strconv.Atoi(strings.TrimSuffix(actionChosenStr, "\n"))
	return actionChosen
}
