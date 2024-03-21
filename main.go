package main

import (
	"fmt"
	"projekt1/fileHandler"
	"projekt1/mySort"
	"sort"
)

func main() {
	fh := fileHandler.FileHandler{FileName: "floaty.txt"}
	fh.ReadFile()
	floatLista := make([]float64, len(fh.GetFloatList()))
	copy(floatLista, fh.GetFloatList())
	fmt.Println(mySort.QuickSort(floatLista, 0, len(floatLista)-1, 1))
	fmt.Println(sort.Float64sAreSorted(floatLista))
}
