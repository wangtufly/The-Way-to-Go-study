package main

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.Map
	scene.Store("A", 97)
	scene.Store("B", 99)
	scene.Store("C", 91)
	scene.Store("D", 17)
	scene.Store("E", 27)

	fmt.Println(scene.Load("E"))

	scene.Delete("E")

	scene.Range(func(key, value interface{}) bool {
		fmt.Println("iterate:", key, value)
		return true
	})
}
