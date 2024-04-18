package submenus

import (
	"fmt"
	"projekt1/utils"
)

func DisplayParams(sortMethodChosen, ssGapChosen, qsPivotChosen, repsChosen, lastTypeChosen int) {
	fmt.Println("Wybrane parametry:")
	fmt.Println("Metoda sortowania: ", utils.SortingMethodNameChooser(sortMethodChosen))
	fmt.Println("Luka dla sortowania Shella: ", utils.SsGapNameChooser(ssGapChosen))
	fmt.Println("Wybór pivota dla sortowania QuickSort: ", utils.PivotNameChooser(qsPivotChosen))
	fmt.Println("Liczba powtórzeń: ", repsChosen)
	fmt.Println("Typ danych: ", utils.TypeNameChooser(lastTypeChosen))
}
