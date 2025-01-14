package main

func main() {
	age, gender := 22, 'm'
	//if age, gender := 22, 'm'; age >= 18 && (gender == 'f' || gender == 70) && true && (true || false) { // 70 is unicode value for 'F'
	if age >= 18 && (gender == 'f' || gender == 70) && true && (true || false) { // 70 is unicode value for 'F'
		println("she is eligible for marriage")
	} else if age >= 21 && (gender == 'M' || gender == 109) && true && (true || false) { // 109 is unicode value for 'm'
		println("he is eligible for marriage")
	} else {
		println("not eligible for marriage")
	}
	//println('F', 'm')//70 109
}
