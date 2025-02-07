package main

import "fmt"

func main() {
	var num1 MyInt1 = 1212
	str1 := num1.ToString()
	fmt.Println(str1)

	var num2 int = 1232

	str2 := MyInt1(num2).ToString()
	fmt.Println(str2)

	var float1 float32 = 123.32

	str3 := MyInt1(float1).ToString()
	fmt.Println(str3)

	// var ok bool = true
	// str := MyInt(ok).ToString() // a bool variable cannot be converted to int and hence it cannot be to MyInt as well

	var (
		n  int     = 100
		n1 MyInt1  = 123
		n2 MyInt2  = 435
		n3 MyInt3  = 124
		f1 float32 = -123.23
	)

	// ToString,Sq,Cube

	str4 := MyInt1(n).ToString()
	s := MyInt2(n).Sq()
	c := MyInt3(n).Cube()

	fmt.Println(str4, s, c)

	str5 := n1.ToString()
	s1 := MyInt2(n1).Sq()
	c1 := MyInt3(n1).Cube()
	fmt.Println(str5, s1, c1)

	str6 := MyInt1(n2).ToString()
	s2 := n2.Sq()
	c2 := MyInt3(n2).Cube()
	fmt.Println(str6, s2, c2)

	str7 := MyInt1(n3).ToString()
	s3 := MyInt2(n3).Sq()
	c3 := n3.Cube()
	fmt.Println(str7, s3, c3)

	str8 := MyInt1(f1).ToString()
	sf1 := MyInt2(f1).Sq()
	sc1 := MyInt3(f1).Cube()

	fmt.Println(str8, sf1, sc1)

}

type MyInt1 int

func (m MyInt1) ToString() string {
	//return strconv.Itoa(int(m))
	return fmt.Sprint(m)
}

type MyInt2 int

func (m MyInt2) Sq() int {
	return int(m * m)
}

type MyInt3 MyInt2

func (m MyInt3) Cube() int {
	return int(m * m * m)
}
