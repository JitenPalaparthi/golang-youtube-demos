package main

func main() {

	var day uint8 = 4

	//switch day := 4; day {
	switch day {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")
	default:
		println("no day")
	}

	num := -149

	switch { // empty switch
	case num >= 0 && num <= 50:
		println(num, " is between and equal to 0 and 50")
	case num > 50 && num <= 100:
		println(num, "is greater than 50 and less or equal to 100")
	case num > 100:
		println(num, "is greater than 100")
	default:
		println(num, "is negative")
	}

	char := 'A'
	switch char {
	case 'a', 'e', 'i', 'o', 'u', 65, 69, 73, 79, 'U': //purposely gave few chars and few utf8 codes
		println(string(char), "is vovel")
	default:
		println(string(char), "is consonent or special character")
	}

}

// 1. There is no mandatory break, break is is by default
// 2. normal switch and empty switch
