package main

import "fmt"

func main() {
	m := map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
	}
	m2 := make(map[string]int) // m2 == empty map
	var m3 map[string]int      // m3 == nil
	fmt.Println(m, m2, m3)

	// 无序 hashmap
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	num, ok := m["1"]
	fmt.Println(num, ok)
	// Zero value
	if num, ok = m["4"]; ok {
		fmt.Println(num)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	num, ok = m["1"]
	fmt.Println(num, ok)

	delete(m, "1")
	num, ok = m["1"]
	fmt.Println(num, ok)

	// 除了 slice, map, function的内建类型都可以作为key
	// struct类型不包含上述字段，也可以作为key
}
