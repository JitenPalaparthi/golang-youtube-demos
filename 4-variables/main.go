package main

import (
	"fmt"
	"unsafe"
)

var G int = 100 // where is this created

const PI float32 = 3.14 // where is this create

func main() {

	var char1 rune = 'A'
	var char2 int32 = 'B'
	var char3 int64 = 'C'
	var char4 int8 = 'D'
	var char5 uint64 = '你' // 20k
	var char6 float64 = '界'
	char1 = char1 + 1

	fmt.Println(char1, char2, char3, char4, char5, char6)
	fmt.Println(string(char1), string(char2), string(char3), string(char4), string(char5), string(int(char6)))
	fmt.Println([]byte(string(char5))) // about

	str1 := "Hello UST Global , Bengalore"         // 28 chars
	str2 := "Hello UST Global, Thiruvananthapuram" // 36 chars
	str3 := "你好, UST Global"                       // 14 chars ?
	var str4 string                                // what is the zero value of it?

	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Println(str4)

	fmt.Println("str1 size", unsafe.Sizeof(str1))
	fmt.Println("str2 size", unsafe.Sizeof(str2))
	fmt.Println("str3 size", unsafe.Sizeof(str3))
	fmt.Println("str4 size", unsafe.Sizeof(str4))
	str4 = "Hello Engineers 🚀" // I am able to mutate?
	// the string is not written to the same memory rather everytime,
	// we try to mutate,a new memory is allocated and assigned

	fmt.Println("str1:", len(str1))
	fmt.Println("str2:", len(str2))
	fmt.Println("str3:", len(str3))
	fmt.Println("str4:", len(str4))

}

// what are chars in Golang?
// All are unicode chars..utf8
// a size of a utf8 chars is from 1-4
// evan a string is nothing but array of bytes -->? will discuss about it later
// is string mutable or immutable?
/* str1 header
--------------
ptr --> pointer to the actual data    8 bytes on 64bit machine
len --> the length of the actual data 8 bytes on 64bit machine
*/

// csp -> communicating sequential processes
// invented null pointer for c++
// inventing null pointer was a billion dollar mistake he has done
