package main

import "os"

func main() {
	//fo := &FileOps{}

	var fo *FileOps
	n, err := fo.WriteToFile("data.txt", []byte("Hello World!"))
	if err != nil {
		println(err.Error())
	} else {
		println(n, "bytes are successfully written to file")
	}

	func() {
		println("End of main func")
	}()

}

type FileOps struct {
	Filename string
}

func (fo *FileOps) WriteToFile(filename string, bytes []byte) (int, error) {
	if fo == nil {
		panic("FileOps variable is nil")
	}

	if fo.Filename == "" {
		panic("invalid file name")
	}
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return -1, err
	}
	defer f.Close()
	return f.Write(bytes)
}
