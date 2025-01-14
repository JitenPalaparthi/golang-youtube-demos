package main

func main() {

	//var age uint8 = 18

	//a, b := 10, 20

	//c := b/a + a*b - 100 + 10 + a/10 - (b / 10 * 20) // This is not an atomic operation
	// there is a precedence
	//c1 := b/a+a*b-100+10+a/10-(b/10*20) > 100

	if age := 18; age >= 18 { // creating a variable and using it for conditinal statement
		println("eligible for vote. age is ", age)
	} else {
		println("not eligible for vote. age is ", age)
	}

}

// arthimetic
// comparision --> true or false
// logical --> true or false
