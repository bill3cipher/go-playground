package main

import "fmt"

type volumer interface {
	Volume() float64
}

type cube struct {
	edge float64
} // edge x edge x edge

type triangularPrism struct {
	base     float64
	attitude float64
	height   float64
} // 0.5 x base x attitude x height

func VolumeOf(v volumer) float64 {
	return v.Volume()
}

func (c cube) Volume() float64 {
	return c.edge * c.edge * c.edge
}
func (t triangularPrism) Volume() float64 {
	return 0.5 * t.base * t.attitude * t.height
}

func main() {
	var a interface{}

	a = 10
	fmt.Printf("%T %v\n", a, a)

	a = "ten"
	fmt.Printf("%T %v\n", a, a)

	a = true
	fmt.Printf("%T %v\n", a, a)

	a = func() string { return "hello" }
	fmt.Printf("%T %v\n", a, a)

	c := &cube{1.23}
	fmt.Printf("%T %v\n", c.Volume(), c.Volume())

	t := &triangularPrism{3, 4, 5}
	fmt.Printf("%T %v\n", t.Volume(), t.Volume())

}
