package main

import (
	"fmt"
	"strings"
)

func main() {
	// Map = (Key,Value)
	// zero value = nil

	// how to declare map
	m := make(map[string]string)
	a := map[string]string{}

	if m == nil {
		fmt.Println("m it is nil")
	}
	if a == nil {
		fmt.Println("a it is nil")
	}

	m["a"] = "apple"
	a["1"] = "banana"

	s := m["a"]
	fmt.Println("m:", s)

	s2 := a["1"]
	fmt.Println("a:", s2)

	count := wordCount("Apple Banana Apple Banana apple")
	fmt.Println(count)
}

func wordCount(s string) map[string]int {
	splitStr := strings.Split(s, " ")
	mapFruits := map[string]int{}
	fmt.Println("splitStr : ", splitStr)
	fmt.Printf("type of splitstr : %T\n", splitStr)
	for i := 0; i < len(splitStr); i++ {
		fmt.Printf("index : %d  value : %s\n", i, splitStr[i])
		fmt.Println(mapFruits)
		mapFruits[splitStr[i]] += 1
	}
	return mapFruits
}
