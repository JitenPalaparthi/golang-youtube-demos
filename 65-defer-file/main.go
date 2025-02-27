package main

import "os"

func main() {
	n, err := WriteToFile("data.txt", []byte("Hello World!"))
	if err != nil {
		println(err.Error())
	} else {
		println(n, "bytes are successfully written to file")
	}

}

func WriteToFile(filename string, bytes []byte) (int, error) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return -1, err
	}
	defer f.Close()
	return f.Write(bytes)
}
