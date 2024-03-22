package main

import (
	"flag"
	"sync"
)

func main() {
	quickRunPtr := flag.Bool("quick-run", true, "if true run default tests")
	sortMethodPtr := flag.String("sort-method", "q", "sorting algorithm(q-quick, i-insertion, h-heap, s-shell), default option: q")
	shellSortGapPtr := flag.Int("ss-gap", 0, "shell sort gap calculation(0-shell, 1-franklazarus), default option: shell")
	quickSortPiwotPtr := flag.Int("qs-pivot", 0, "choose pivot selection method for quicksort(0-high, 1-low, 2-mid, 3-rand), default option: high")
	variableTypePtr := flag.Int("var-type", 0, "variable type for sorting (0-int, 1-float), default option: int")
	inputFilePtr := flag.String("input", "input.txt", "name of input file, default: input.txt")
	repetitionPtr := flag.Int("reps", 1, "number of repetitions, default: 1")
	generationLengthPtr := flag.Int("generate-length", 10000, "number of lines in new input, default: 10000")
	generationPtr := flag.Bool("generate-new-input", false, "generate new input file, may overwrite existing one if filename option is name of existing file, default option: false")
	outputFilePtr := flag.String("output", "output.csv", "name of output file, by default: output.csv")
	flag.Parse()
	if *quickRunPtr {
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			quickRun()
			wg.Done()
		}()
		wg.Wait()
	} else {
		run(*inputFilePtr, *outputFilePtr, *sortMethodPtr, *shellSortGapPtr, *quickSortPiwotPtr, *variableTypePtr, *repetitionPtr, *generationLengthPtr, *generationPtr)
	}
}
