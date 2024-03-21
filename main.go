package main

import (
	"fmt"
	"projekt1/sort"
)

func main() {
	var lista []int
	lista = append(lista, 3, 6, 1, 4, 7, 9, 11, 4, 6, 1, 2, 4, 2, 5, 9)
	fmt.Println(lista)
	fmt.Println(sort.InsertionSort(lista))
}
