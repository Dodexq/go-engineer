package main

import "fmt"

func main() {
	type Point struct {
		X int
		Y int
	}
	p := Point{
		X: 5,
		Y: 3,
	}
	fmt.Println(p)

	p = Point{X: 12, Y: 25}
	fmt.Println(p)

	x := Point{
		Y: 25,
		X: 24,
	}
	fmt.Println(x)

	fmt.Println(p.X)
	fmt.Println(x.Y)

	p.X = 10
	fmt.Println(p.X)
}
