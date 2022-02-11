package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	m2 := make(map[string]int) //m2=empty map

	var m3 map[string]int // m3=nil

	fmt.Println(m)
	fmt.Println(m2)
	fmt.Println(m3)
	fmt.Println("traversing map:")
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("getting values:")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	if causeName, ok := m["caurse"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key does not exist")
	}
	fmt.Println("delete values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
