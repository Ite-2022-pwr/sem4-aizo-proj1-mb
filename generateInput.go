package main

import (
	"math/rand"
	"os"
	"strconv"
)

func generateInput(lines, varType int, filename string) {
	f, err := os.Create(filename)
	Check(err)

	defer func(f *os.File) {
		err := f.Close()
		Check(err)
	}(f)
	if varType == 0 {
		for i := 0; i < lines; i++ {
			temp := strconv.Itoa(rand.Int() % 1000)
			_, err := f.WriteString(temp + "\n")
			Check(err)
		}
	} else if varType == 1 {
		for i := 0; i < lines; i++ {
			temp := strconv.FormatFloat(rand.Float64()*10, 'g', -1, 64)
			_, err := f.WriteString(temp + "\n")
			Check(err)
		}
	}

	err = f.Sync()
	if err != nil {
		Check(err)
	}

}
