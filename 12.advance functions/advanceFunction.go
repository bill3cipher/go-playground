package main

import "fmt"

func main() {
	//Higher Order Function
	s := HOFGreeting(func() string {
		return "New"
	})

	//or
	nameFn := func() string { return "New2" }
	s2 := HOFGreeting_2(nameFn)

	fmt.Println(s)
	fmt.Println(s2)

	fmt.Println(HOFGreeting_2(newNameFunc("NewZa")))

	//Closure function
	counter := newCounterFunc()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

// function as parameters
func HOFGreeting(nameFn func() string) string {
	return fmt.Sprintf("hello, %s", nameFn())
}

//function as returns
func newNameFunc(name string) nameFunc {
	return func() string {
		return name
	}
}

// function as parameters
type nameFunc func() string

func HOFGreeting_2(nameFn nameFunc) string {
	return fmt.Sprintf("hello, %s", nameFn())
}

//Closure function
func newCounterFunc() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}
