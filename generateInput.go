package main

import (
	"math/rand"
	"os"
	"strconv"
)

func generateInput(lines int, filename string) {
	f, err := os.Create(filename)
	check(err)

	defer func(f *os.File) {
		err := f.Close()
		check(err)
	}(f)

	for i := 0; i < lines; i++ {
		temp := strconv.Itoa(rand.Int() % 1000)
		_, err := f.WriteString(temp + "\n")
		check(err)
	}

	err = f.Sync()
	if err != nil {
		check(err)
	}

}
