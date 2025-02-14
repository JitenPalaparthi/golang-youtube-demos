package main

import "fmt"

func main() {

	var num int = 100

	var ok bool = true

	var str = "Hello World"

	var iempty IEmpty = num // exactly like any

	fmt.Println(iempty)
	iempty = ok
	fmt.Println(iempty)
	iempty = str
	fmt.Println(iempty)

	str1, ok1 := iempty.(string)
	if ok1 {
		fmt.Println(str1, ok1)
	} else {
		fmt.Println("Could not assert it to string")
	}

	iempty = &Rect{10.2, 12.45}

	r1, ok1 := iempty.(*Rect)
	if ok1 {
		fmt.Println(r1)
	} else {
		fmt.Println("could not type assert to *Rect")
	}

}

type IEmpty interface {
}

type Rect struct {
	L, B float32
}
