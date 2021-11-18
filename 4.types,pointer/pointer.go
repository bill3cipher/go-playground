package main

import "fmt"

func main() {
	var s string
	var p *string

	fmt.Println("s: ", s)
	fmt.Println("p: ", p)
	fmt.Println("&s: ", &s)
	p = &s //address of s
	fmt.Println("p: ", p)
	fmt.Println("*p: ", *p)
	*p = "hello world"
	fmt.Println("*p with text: ", *p)

	n := 2
	double(&n)
	fmt.Println(n)

	name := "new"
	appendGreeting(&name)
	fmt.Println(name)
}

func double(d *int) {
	*d = *d * 2
}

func appendGreeting(s *string) {
	*s = "Hi, " + *s
}
