package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s1 S1
	var s2 S2

	a1 := s1.Area(10.12, 12.34)
	a2 := s2.Area(12.46)

	fmt.Println("Area of S1:", a1)
	fmt.Println("Area of S2:", a2)

	fmt.Println("Size of s1:", unsafe.Sizeof(s1))
	fmt.Println("Size of s2:", unsafe.Sizeof(s2))

	var s3 *S1 = &S1{}
	var s4 *S2 = &S2{}
	s5 := new(S1)

	a3 := s3.Area(10.12, 12.34)
	a4 := s4.Area(12.4)
	a5 := s5.Area(12.3, 12.45)
	fmt.Println("Area of S1 using s3:", a3)
	fmt.Println("Area of S2 using s4:", a4)
	fmt.Println("Area of S1 using s5:", a5)

	fmt.Printf("Address of s1: %p\n", &s1)
	fmt.Printf("Address of s2: %p\n", &s2)
	fmt.Printf("Address of s3: %p\n", s3)
	fmt.Printf("Address of s4: %p\n", s4)
	fmt.Printf("Address of s5: %p\n", s5)

	slice1 := []int{} // Ptr,Len,Cap
	slice2 := []int{}
	slice3 := make([]int, 0)

	ptr1 := (*[3]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&slice1))))
	fmt.Printf("Ptr in slice1:0x%x\n", ptr1[0])
	fmt.Println("Len in slice1:", ptr1[1])
	fmt.Println("Cap in slice1:", ptr1[2])

	ptr2 := (*[3]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&slice2))))
	fmt.Printf("Ptr in slice2:0x%x\n", ptr2[0])
	fmt.Println("Len in slice2:", ptr2[1])
	fmt.Println("Cap in slice2:", ptr2[2])

	ptr3 := (*[3]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&slice3))))
	fmt.Printf("Ptr in slice3:0x%x\n", ptr3[0])
	fmt.Println("Len in slice3:", ptr3[1])
	fmt.Println("Cap in slice3:", ptr3[2])

}

type S1 struct{}
type S2 struct{}

func (s *S1) Area(l, b float32) float32 {
	return l * b
}

func (s *S2) Area(side float32) float32 {
	return side * side
}
