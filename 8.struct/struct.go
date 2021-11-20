package main

import "fmt"

type rectangle struct {
	w float64
	l float64
}

func main() {
	rec := rectangle{
		w: 4,
		l: 5,
	}

	rec.w = 6
	rec.l = 7

	fmt.Println(rec.w)
	fmt.Println(rec.l)
}
