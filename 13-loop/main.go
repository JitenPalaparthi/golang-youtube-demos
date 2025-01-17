package main

func main() {

	// for loop can have multiple variables

	for i, j := 1, 10; i <= j; i, j = i+1, j-1 {
		println(i, j)
	}

	// there is nested loops
	println("nested loops")
	// done := false
	// for i := 1; i <= 10; i++ {
	// 	if done {
	// 		break
	// 	}
	// 	for j := 5; j <= 10; j++ {
	// 		println("i:", i, "j:", j)
	// 		if i == j {
	// 			done = true
	// 			break // break breaks only the inner loop but not the outer loop by default
	// 		}
	// 	}
	// }

out:
	for i := 1; i <= 10; i++ {
		for j := 5; j <= 10; j++ {
			println("i:", i, "j:", j)
			if i == j {
				break out // break not only breaks the inner loop but breaks to the lable
			}
		}
	}
}
