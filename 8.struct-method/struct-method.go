package main

import "fmt"

type rectangle struct {
	w float64
	l float64
}

type Book struct {
	Name   string
	Author string
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

	fmt.Println(Area(rec))
	fmt.Println(rec.Area())

	book := Book{
		Name:   "Harry Potter",
		Author: "J. K. Rowling",
	}

	fmt.Println(book.String())
	book.SetName("test")
	fmt.Println(book.String())
}

// This is Function
func Area(rec rectangle) float64 {
	return rec.w * rec.l
}

// This is Method
func (rec rectangle) Area() float64 {
	return rec.w * rec.l
}

func (book Book) String() string {
	return book.Name + " by " + book.Author
}

func (book *Book) SetName(name string) {
	book.Name = name
}
