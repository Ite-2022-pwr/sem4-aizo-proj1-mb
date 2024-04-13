package utils

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"projekt1/timeTrack"
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

func GenerateRandomIntList(length int) []int {
	output := make([]int, length)
	for i := 0; i < length; i++ {
		output[i] = rand.Int() % 1000
	}
	return output
}

func GenerateRandomFloat64List(length int) []float64 {
	output := make([]float64, length)
	for i := 0; i < length; i++ {
		output[i] = rand.Float64() * 1000
	}
	return output
}

func GenerateRandomFloat32List(length int) []float32 {
	output := make([]float32, length)
	for i := 0; i < length; i++ {
		output[i] = float32(rand.Float32() * 1000)
	}
	return output
}

func GenerateRandomInt32List(length int) []int32 {
	output := make([]int32, length)
	for i := 0; i < length; i++ {
		output[i] = int32(rand.Int() % 1000)
	}
	return output
}

func GenerateRandomInt64List(length int) []int64 {
	output := make([]int64, length)
	for i := 0; i < length; i++ {
		output[i] = int64(rand.Int() % 1000)
	}
	return output
}

func GenerateRandomStringList(length int) []string {
	output := make([]string, length)
	charRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	for i := 0; i < length; i++ {
		b := make([]rune, (rand.Int()%16)+1)
		for i := range b {
			b[i] = charRunes[rand.Intn(len(charRunes))]
		}
		tempStr := string(b)
		output[i] = tempStr
	}
	return output
}
