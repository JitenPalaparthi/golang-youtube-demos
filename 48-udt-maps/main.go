package main

import (
	"errors"
	"fmt"
)

func main() {
	var m1 MyMap = make(MyMap) // declared and instantiated the map
	m1["560086"] = "Bangalore-1"
	m1["560096"] = "Bangalore-2"
	m1["560034"] = "Bangalore-3"
	m1["560006"] = "Bangalore-4"
	m1["ok"] = false

	if keys, err := m1.GetKeys(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Keys:", keys)
	}

	if values, err := m1.GetValues(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Values:", values)
	}

	m1.Display()

	fmt.Println("Delete method")

	if err := m1.Delete("ok"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("successfully deleted")
	}
	println()
	fmt.Println("Display after Delete method")
	m1.Display()

	if err := m1.Delete("ok"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("successfully deleted")
	}

	println()
	println("Calling methods of MyMap on normal map by type casting")
	var m2 map[string]any = make(map[string]any)

	m2["560086"] = "Bangalore-1"
	m2["560096"] = "Bangalore-2"
	m2["560034"] = "Bangalore-3"
	m2["560006"] = "Bangalore-4"
	m2["ok"] = false

	MyMap(m2).Display()
	if keys, err := MyMap(m2).GetKeys(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Keys:", keys)
	}
	if values, err := MyMap(m2).GetValues(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Values:", values)
	}

	//m3 := make(map[string]string)

	//MyMap(m3). // cannot cast this to MyMap as both then underlining types are different

}

type MyMap map[string]any

func (m MyMap) GetKeys() ([]string, error) {
	if m == nil {
		return nil, errors.New("nil map")
	}
	keys := make([]string, len(m))

	i := 0

	for k := range m { // range loop on map gives key and value for iteration
		keys[i] = k
		i++
	}
	return keys, nil
}

func (m MyMap) GetValues() ([]any, error) {
	if m == nil {
		return nil, errors.New("nil map")
	}
	var values []any
	for _, v := range m {
		values = append(values, v) // if the slice is nil , append intantiate the slice
	}
	return values, nil
}

func (m MyMap) Delete(key string) error {
	if m == nil {
		return errors.New("nil map")
	}
	if _, ok := m[key]; !ok {
		return errors.New("key does not exist")
	}
	delete(m, key)
	return nil
}

func (m MyMap) Display() error {
	if m == nil {
		return errors.New("nil map")
	}
	for k, v := range m {
		fmt.Println("Key:", k, "Value:", v)
	}
	return nil
}

// type StringSlice []string
//
//func (ss StringSlice) Sort() (StringSlice, error) {

//}

//type MyMap1 map[string]StringSlice

// GetKeys
// GetValues
// Delete
// Sort. the Sort should work on the value of MyMap1
