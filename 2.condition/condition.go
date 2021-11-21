package main

import (
	"fmt"
	"strconv"
)

func main() {
	name := "new"
	if name != "" {
		fmt.Println("Hello, ", name)
	} else {
		fmt.Println("Who am i?")
	}

	//short statement
	num := "10"
	if output, err := strconv.Atoi(num); err != nil {
		fmt.Println("NaN: ", err)
	} else {
		fmt.Println(output)
	}

	//full statement
	out, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("NaN: ", err)
	} else {
		fmt.Println(out)
	}

	isOdd := isOdd(0)
	fmt.Println(isOdd)

}

func isOdd(n int) bool {
	if n%2 == 0 {
		return false
	} else {
		return true
	}
}
