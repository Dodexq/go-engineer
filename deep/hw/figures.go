package main

import "fmt"

type Rectangle struct{}
type Circle struct{}

func (r *Rectangle) Area(a float64, b float64) float64 {
	return a * b
}

func (r *Rectangle) Type() {
	fmt.Println("Прямоугольник")
}

func (c *Circle) Area(r float64) float64 {
	return 3.14 * r * r
}

func (c *Circle) Type() {
	fmt.Println("Круг")
}

func main() {
	r := &Rectangle{}
	fmt.Println(r.Area(2, 3))
	r.Type()

	fmt.Println()

	c := &Circle{}
	fmt.Println(c.Area(4))
	c.Type()
}
