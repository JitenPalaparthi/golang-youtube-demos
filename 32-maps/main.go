package main

import "fmt"

func main() {
	mymap := make(map[string]any)
	mymap["560086"] = "Bengalore-1"
	mymap["560096"] = "Bengalore-2"
	mymap["560034"] = "Bengalore-3"
	mymap["695011"] = "Trivandrum Medical college"
	mymap["695012"] = 2
	v := mymap["560086"]
	fmt.Println(v)

	v, ok := mymap["695013"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("no key and value pair")
	}
}
