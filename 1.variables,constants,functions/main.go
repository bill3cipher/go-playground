package main

import "fmt"

func main() {
	//variable
	var a int
	var b string = "new, nue, nil"
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//infer variable
	var i = 1
	var j = false
	k := "This is string?"
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	greeting(b)
	greeting_return(k)
	greetingWithAge("new", 25)
	greetingWithAgeAndDrink("new", 25, "coke zero")
	swap(i, 2)
}

//function with parameter no return
func greeting(name string) {
	fmt.Println("Hello, ", name)
}

//return function with parameter
func greeting_return(name string) string {
	return fmt.Sprintf("Hello_Return, %s", name)
}

//return function with 2 parameters
// greetingWithAge("Pallat", 30) should return "Hello, Pallat. You are 30 years old."
func greetingWithAge(name string, age uint) string {
	return fmt.Sprintf("Hello, %s. You are %d years old", name, age)
}

//return function with 3 parameters
// greetingWithAgeAndDrink("Pallat", 30, "Cola") should return "Hello, Pallat. You are 30 years old and your favorite drink is Cola."
func greetingWithAgeAndDrink(name string, age uint, drink string) string {
	return fmt.Sprintf("Hello, %s. You are %d years old and your favorite drink is %s.", name, age, drink)
}

//function return 2 values
func swap(a, b int) (int, int) {
	return b, a
}
