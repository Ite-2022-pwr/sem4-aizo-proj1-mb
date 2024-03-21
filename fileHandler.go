package main

type fileHandler[T any] struct {
	path  string
	lines int
	list  []T
}

func newFileHandler[T any](path string) *fileHandler[T] {
	fh := fileHandler[T]{path: path}
	fh.lines = 0
	return &fh
}

func (fh fileHandler[T]) readFile() {

}
