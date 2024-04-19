package utils

import (
	"fmt"
	"os"
	"projekt1/utils/listGenerators"
)

func SaveListToFile[T any](list []T, filename string) {
	f, err := os.Create(filename)
	Check(err)
	defer func(f *os.File) {
		err := f.Close()
		Check(err)
	}(f)
	_, err = f.WriteString(fmt.Sprintf("%d\n", len(list)))
	for _, v := range list {
		_, err := f.WriteString(fmt.Sprintf("%v\n", v))
		Check(err)
	}
	err = f.Sync()
	if err != nil {
		Check(err)
	}
}

func GenerateList(typeChosen, generationStrategy, length int) (list []any) {
	switch typeChosen {
	case 0:
		list = CastToAnySlice(listGenerators.GenerateStdIntList(generationStrategy, length))
	case 1:
		list = CastToAnySlice(listGenerators.GenerateFloat64List(generationStrategy, length))
	case 2:
		list = CastToAnySlice(listGenerators.GenerateFloat32List(generationStrategy, length))
	case 3:
		list = CastToAnySlice(listGenerators.GenerateInt32List(generationStrategy, length))
	case 4:
		list = CastToAnySlice(listGenerators.GenerateInt64List(generationStrategy, length))
	case 5:
		list = CastToAnySlice(listGenerators.GenerateStringList(generationStrategy, length))
	default:
		list = CastToAnySlice(listGenerators.GenerateStdIntList(generationStrategy, length))
	}
	return list
}
