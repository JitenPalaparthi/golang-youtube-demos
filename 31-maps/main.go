package main

import "fmt"

func main() {

	var mymap map[string]any // only declaration but not instantiation
	// var mymap2 map[[]int]any

	// var str1, str2 string = "Hello", "World"
	// if str1 == str2 {
	// 	println("both strings are same")
	// }
	// slice1 := []int{10, 20, 30}
	// slice2 := []int{10, 20, 30}
	// // if slice1 == slice2 { // equal == cannot be applied on slices

	// // }
	// if slices.Equal(slice1, slice2) {
	// 	println("yes both are equal")
	// }

	if mymap == nil {
		fmt.Println("nil map")
		mymap = make(map[string]any) // make is used to instantiate the map
	}

	mymap["560086"] = "Bengalore-1"
	mymap["560096"] = "Bengalore-2"
	mymap["560034"] = "Bengalore-3"
	mymap["695011"] = "Trivandrum Medical college"
	mymap["695012"] = 2

	// var mymap2 map[string]string = map[string]string{"560086": "Bangalore-1", "560096": "Bangalore-2"}

	// slice1 := make([]int, 10)
	// slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for k, v := range mymap {
		fmt.Println("Key:", k, "Values:", v)
	}

}

// Hash Table
// O(1) or  O(n) maps are efficient
// maps contain keys and values, values are mapped to the keyes
// What is Hash -->

// SHA1 (hello1.txt) = 0a4d55a8d778e5022fab 701977c5d840bbc486d0
// SHA1 (hello2.txt) = 0a4d55a8d778e5022fab701977c5d840bbc486d0
// SHA1 (hello2.txt) = 053bebe714723f6bd76ba08f869ad8113a408d3e

// 1. keys are values are pairs, values are mapped to the keys
// 1.1 a map can be nil, if not instantiated
// 2. what datatype can be the key? Any type on which one can apply == operator can be the key
// 3. When a map is created , based on the size of the map, buckets are created
// 4. Each bucket has keys and values ( parallel arrays)
/* 5.
Bucket-0
---------
i  K  V
0 |	 |	|
1 |	 |	|
2 |	 |	|
3 |	 |	|
4 |	 |	| --> linked list
5 |	 |	|
6 |	 |	|
7 |	 |	|

Bucket-1
---------
i  K  V
0 |	 |	|
1 |	 |	|
2 |	 |	|
3 |	 |	|
4 |	 |	|
5 |	 |	|
6 |	 |	|
7 |	 |	|


Bucket-2
---------
i  K  V
0 |	 |	|
1 |	 |	|
2 |	 |	|
3 |	 |	|
4 |	 |	|
5 |	 |	|
6 |	 |	|
7 |	 |	|
How many buckets are created?
Based on the load factor. number of elements / buckets --> 6.5
If there are 14 elemens(kv pairs) and 2 buckets the Load factor is 7.0
When ever the load factor goes beyond 6.5, the the nuber of buckets are doubled
When the same hash values are generated for more then one key, a collision happen
When the collision happens a linked list is created and attached to the bucket
*/
// 6. map is not ordered
// 7. map is not thread safe(Goroutine safe), if multiple go routines try to access the map, there is a chance of race conditions
