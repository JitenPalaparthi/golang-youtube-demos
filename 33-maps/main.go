package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {

	mymap := make(map[string]any)

	mymap["cab-1"] = 345.50
	mymap["coffee-1"] = int8(40)
	mymap["airport-travel-1"] = float32(600.50)
	mymap["in-flight-meal-1"] = int16(340)
	mymap["misc-1"] = float32(123.25)
	mymap["misc-2"] = "12312.123" // this is not taken care
	total, err := TotalExpense(mymap)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Total Expenses:", total)
	}

	delete(mymap, "misc-2")
	fmt.Println(mymap)
	delete(mymap, "misc-3")
	fmt.Println(mymap)
	var mymap2 map[string]any // nil map
	if mymap2 == nil {
		fmt.Println("nil map")
	}
	fmt.Println(mymap2)
	delete(mymap2, "abcd")
	err = DeleteElem(mymap, "abcd")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("key-value pair successfully deleted from the map")
	}
}

func TotalExpense(mymap map[string]any) (float32, error) {
	if mymap == nil {
		return 0, errors.New("input map is nil")
	}
	sum := float32(0)
	for _, v := range mymap {
		if IsNumber(v) {
			switch s := v.(type) {
			case int:
				sum += float32(s)
			case uint8:
				sum += float32(v.(uint8))
			case uint16:
				sum += float32(s)
			case uint32:
				sum += float32(s)
			// case uint64, uint, int8, int16, int32, int64:
			// 	sum += float32(s) // This does not work bcz s value does not know what type it belongs to among more type
			case uint64:
				sum += float32(s)
			case int8:
				sum += float32(s)
			case int16:
				sum += float32(s)
			case int32:
				sum += float32(s)
			case int64:
				sum += float32(s)
			case float32:
				sum += s
			case float64:
				sum += float32(s)
			case string:
				v1, err := strconv.ParseFloat(s, 64)
				if err == nil {
					sum += float32(v1)
				} else {
					v2, err := strconv.Atoi(s)
					if err == nil {
						sum += float32(v2)
					}
				}
			}
		}
	}
	return sum, nil
}

func IsNumber(v any) bool {
	switch v.(type) { // can take the value into a variable or directly type assert as and when required
	case int, uint, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64:
		return true
	case string:
		_, err1 := strconv.ParseFloat(v.(string), 64)
		_, err2 := strconv.Atoi(v.(string))
		if err1 == nil || err2 == nil {
			return true
		} else {
			return false
		}

	default:
		return false
	}
}

func DeleteElem(mymap map[string]any, key string) error {
	if mymap == nil {
		return errors.New("input map is nil")
	}
	_, ok := mymap[key]
	if !ok {
		return errors.New("key does not exist")
	}
	delete(mymap, key)
	return nil
}
