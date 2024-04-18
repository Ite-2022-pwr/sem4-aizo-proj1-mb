package menu

import (
	"fmt"
	"projekt1/menu/submenus"
	"projekt1/runner"
)

func Menu() {
	var lastList []any
	var lastTypeChosen int
	var sortMethodChosen = 0
	var ssGapChosen = 0
	var qsPivotChosen = 0
	var repsChosen = 1
	var lastRunTimes []int64
	var lastRunLines []int

	for true {
		fmt.Println()
		actionChosen := submenus.ChooseAction()
		switch actionChosen {
		case 0:
			lastList, lastTypeChosen = submenus.InputDataFromFile()
		case 1:
			lastList, lastTypeChosen = submenus.GenerateData()
		case 2:
			submenus.OutputListToFile(lastList)
		case 3:
			lastRunLines, lastRunTimes = runner.Run(sortMethodChosen, lastTypeChosen, ssGapChosen, qsPivotChosen, repsChosen, lastList)
		case 4:
			submenus.DisplayList(lastList)
		case 5:
			sortMethodChosen, ssGapChosen, qsPivotChosen = submenus.ChooseSorting()
		case 6:
			repsChosen = submenus.ChooseReps()
		case 7:
			submenus.SaveResultToFile(lastRunLines, lastRunTimes)
		case 8:
			submenus.DisplayParams(sortMethodChosen, ssGapChosen, qsPivotChosen, repsChosen, lastTypeChosen)
		case 9:
			return
		}
	}
}
