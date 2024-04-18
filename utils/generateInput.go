package utils

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"projekt1/timeTrack"
	"projekt1/utils/listGenerators"
	"strconv"
	"time"
)

func GenerateInputToFile(lines, varType int, filename string) {
	name := fmt.Sprintf("Generating %d lines of %s", lines, filename)
	startTime := time.Now()
	log.Printf("%s started at: %s", name, startTime)
	defer timeTrack.TimeTrack(startTime, name)
	f, err := os.Create(filename)
	Check(err)
	_, err = f.WriteString(fmt.Sprint(lines) + "\n")
	Check(err)

	defer func(f *os.File) {
		err := f.Close()
		Check(err)
	}(f)
	if varType == 0 {
		for i := 0; i < lines; i++ {
			temp := strconv.Itoa(rand.Int() % 100000)
			_, err := f.WriteString(temp + "\n")
			Check(err)
		}
	} else if varType == 1 {
		for i := 0; i < lines; i++ {
			temp := strconv.FormatFloat(rand.Float64()*100000, 'g', -1, 64)
			_, err := f.WriteString(temp + "\n")
			Check(err)
		}
	}

	err = f.Sync()
	if err != nil {
		Check(err)
	}
}

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
