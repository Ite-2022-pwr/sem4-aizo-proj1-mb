package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"projekt1/timeTrack"
	"strconv"
	"time"
)

func generateInputToFile(lines, varType int, filename string) {
	name := fmt.Sprintf("Generating %d lines of %s", lines, filename)
	startTime := time.Now()
	log.Printf("%s started at: %s", name, startTime)
	defer timeTrack.TimeTrack(startTime, name)
	f, err := os.Create(filename)
	Check(err)
	_, err = f.WriteString(string(lines) + "\n")
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
			temp := strconv.FormatFloat(rand.Float64()*1000, 'g', -1, 64)
			_, err := f.WriteString(temp + "\n")
			Check(err)
		}
	}

	err = f.Sync()
	if err != nil {
		Check(err)
	}

}
