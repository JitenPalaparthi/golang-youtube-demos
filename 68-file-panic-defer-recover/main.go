package main

import (
	"fmt"
	"os"
)

func main() {
	//fo := &FileOps{}
	defer func() {
		println("End of main func")
	}()
	defer println("End of main --1")
	defer RecoverMe()
	var fo *FileOps = &FileOps{}
	n, err := fo.WriteToFile([]byte("Hello World!"))
	if err != nil {
		println(err.Error())
	} else {
		println(n, "bytes are successfully written to file")
	}
	n1 := 0
	for i := 1; i <= 20; i++ {
		if i%n1 == 0 {
			println(i)
		}
	}

	defer func() { // This is not executed since called after panic
		println("deferred call after panic")
	}()

}

type FileOps struct {
	Filename string
}

type FileErr struct {
	Code    int
	FileOps *FileOps
	Message string
}

func (fo *FileOps) WriteToFile(bytes []byte) (int, error) {
	defer func() {
		println("End of WriteToFile function")
	}()

	defer RecoverMe()

	// defer func() {
	// if r := recover(); r != nil {
	//fmt.Println(r)
	// switch v := r.(type) {
	// case *FileErr:
	// fmt.Println("Panic Code:", v.Code)
	// fmt.Println("Panic Message:", r.(*FileErr).Message) // can also assrt on r.
	// fmt.Println("Panic Object", v.FileOps)              // can also use v
	// }

	// }
	// }()

	if fo == nil {
		panic(&FileErr{Code: 101, FileOps: nil, Message: "FileOps variable is nil"})
		//	panic("FileOps variable is nil")
	}

	if fo.Filename == "" {
		panic(&FileErr{Code: 102, FileOps: nil, Message: "invalid file name"})
		//panic("invalid file name")
	}
	f, err := os.OpenFile(fo.Filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return -1, err
	}
	defer f.Close()
	return f.Write(bytes)

}

func RecoverMe() {
	if r := recover(); r != nil {
		//fmt.Println(r)
		//	fmt.Println("What type", reflect.TypeOf(r))
		switch v := r.(type) {
		case *FileErr:
			fmt.Println("Panic Code:", v.Code)
			fmt.Println("Panic Message:", r.(*FileErr).Message) // can also assrt on r.
			fmt.Println("Panic Object", v.FileOps)              // can also use v
		default:
			fmt.Println(r)
		}
	}

	// if there is no panic.. r is nil .. and hence does nothing
}
