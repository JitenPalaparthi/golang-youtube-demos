package main

func main() {

	num := 48

	switch { // empty switch
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, "is divisible by 2")
	}

	println("\nfalse negative?...")
	num = 2
	// false negative
	switch { // empty switch
	case num%2 == 0:
		println(num, "is divisible by 2")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%8 == 0:
		println(num, "is divisible by 8")
	}

	num = 11

	switch {
	case num >= 10:
		println(num, "is greater than or equal 10")
		fallthrough
	case num > 20:
		println(num, "is greatr than 20")
	case num <= 0:
		println(num, "is less than or equal 0")
	}

}

// what if the same effect of removing break in other programming languages in Go?
// we use a speacial keyword called fallthrough
// here is the catch, if the fallthrough is not properly used, it gives false negative results
