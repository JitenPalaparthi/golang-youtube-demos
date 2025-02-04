package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"unsafe"
)

func main() {
	str1 := "Hello World"
	str2 := "你好，世界"

	for i := 0; i < len(str1); i++ {
		fmt.Print(string(str1[i]))
	}

	println()

	for i := 0; i < len(str2); i++ {
		fmt.Print(string(str2[i]))
	}
	println()
	for _, v := range str2 {
		print(string(v))
	}
	println()
	println("Converting chinees string to rune slice")
	bytes := []byte(str2)
	chars := []rune(str2)
	chars2 := []rune(str1)

	fmt.Println("Len of bytes:", len(bytes), "size of bytes:", unsafe.Sizeof(bytes))
	fmt.Println("Len of chars:", len(chars), "size of chars:", unsafe.Sizeof(chars))
	fmt.Println("Len of chars2:", len(chars2), "size of chars2:", unsafe.Sizeof(chars2))

	for i := 0; i < len(bytes); i++ {
		fmt.Print(string(bytes[i]))
	}
	println()
	for i := 0; i < len(chars); i++ {

		fmt.Print("size of-->", unsafe.Sizeof(chars[i]), "-->", string(chars[i]))
	}

	println()

	for i := 0; i < len(chars2); i++ {
		fmt.Print(string(chars2[i]))
	}
	println()
	println("Hashcode(sha1) of str1")

	hash := sha1.New()
	hash.Write([]byte(str1))
	sha1hash := hex.EncodeToString(hash.Sum(nil))
	fmt.Println(sha1hash)

	hash256 := sha256.New()
	hash256.Write([]byte(str1))
	sha256hash := hex.EncodeToString(hash256.Sum(nil))
	fmt.Println(sha256hash)

}
