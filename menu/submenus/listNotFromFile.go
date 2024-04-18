package submenus

import "projekt1/utils"

func GenerateData() (list []any, typeChosen int) {
	typeChosen = ChooseType()
	generationStrat := ChooseGenerationStrategy()
	length := EnterLength()
	list = utils.GenerateList(typeChosen, generationStrat, length)
	return list, typeChosen
}
