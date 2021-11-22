package main

import (
	"fmt"

	"github.com/chakhrits/goplayground/11.pack/book"
)

func main() {
	b := book.New()

	fmt.Printf("%T %v\n", b, b)

	b.Name = "The Lord Of The Ring"
	fmt.Printf("%T %v\n", b, b)
}
