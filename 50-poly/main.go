package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//var p1 *int // p1 also has some address.. address is assigned to p1

	var num int
	r1 := new(R) //
	s1 := new(S) //

	var (
		r2 R
		s2 S
		t1 struct{}
		t2 T
	)

	fmt.Printf("%p size of r1:%d\n", &r1, unsafe.Sizeof(r1))
	fmt.Printf("%p size of s1:%d\n", &s1, unsafe.Sizeof(s1))

	fmt.Printf("%p size of r1:%d\n", r1, unsafe.Sizeof(*r1))
	fmt.Printf("%p size of s1:%d\n", s1, unsafe.Sizeof(*s1))

	fmt.Printf("%p size of r2:%d\n", &r2, unsafe.Sizeof(r2))
	fmt.Printf("%p size of s2:%d\n", &s2, unsafe.Sizeof(s2))
	fmt.Printf("%p size of t1:%d\n", &t1, unsafe.Sizeof(t1))
	fmt.Printf("%p size of num:%d\n", &num, unsafe.Sizeof(num))
	fmt.Printf("%p size of num:%d\n", &t2, unsafe.Sizeof(t2))

}

type R struct{} // size is zero
type S struct{} // size is zero
// type B byte     // it still allocates 1 bytes
type T struct{ A int }

func (r *R) Area(l, b float32) float32 {
	return l * b
}

func (s *S) Area(side float32) float32 {
	return side * side
}
