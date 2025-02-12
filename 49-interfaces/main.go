package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "Hello World")

	fw := NewFileWriter("data.txt")

	fmt.Fprintln(fw, "Hello World")

	fmt.Fprintln(fw, "Hello Golang Minds. Happy learning!")
}

// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }

type FileWriter struct {
	FileName string
}

func NewFileWriter(fileName string) *FileWriter {
	return &FileWriter{FileName: fileName}
}

func (fw *FileWriter) Write(p []byte) (n int, err error) {
	if fw == nil {
		return 0, errors.New("nil File Writer")
	}
	file, err := os.OpenFile(fw.FileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		return 0, err
	}
	defer file.Close() // this can be discussed later

	return file.Write(p)
}

// interface is a contract, specification
// interface is a shared behaviour
// inteface contains collection of method signatures of a type that implment.
// But it does not contain how methods to be implemented

// FileModes

// 0 6 4 4
// 0 7 7 7

// Read   	-- 4
// Write 	-- 2
// Execute 	-- 1
