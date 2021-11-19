package main

import "fmt"

func main() {
	// array is immutable

	var num [4]int

	num[0] = 1
	num[1] = 2

	//slice

	var s []int
	if s == nil {
		fmt.Println("it's nil")
	}

	a := [...]int{1, 3, 5, 7, 9}
	s = a[:]
	fmt.Printf("%d %d %p %p\n", len(s), cap(s), s, &a)
	s = append(s, 11, 13)
	fmt.Printf("%d %d %p %p\n", len(s), cap(s), s, &a)
	abc()
	efg()
	cde()
}

func abc() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}
	s = s[:3]
	fmt.Print(s)
	// [apple banana coconut]
}

func efg() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}
	s = s[4:7]
	fmt.Print(s)
	// * [elderberries figs grapes]
}

func cde() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}
	s = s[2:5]
	fmt.Print(s)
	// * [coconut durian elderberries]
}
