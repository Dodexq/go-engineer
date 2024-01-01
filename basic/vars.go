package main

import "fmt"

func main() {
	var v1 int
	fmt.Println(v1)

	var v2 int = 220
	fmt.Println(v2)

	v3 := 330
	fmt.Println(v3)

	var v5, v6 = 1, 2
	fmt.Println(v5, v6)

	var (
		v01 = 1
		v02 = "string"
	)
	fmt.Println(v01, v02)

	const EarthRadius int = 6371000
	_ = EarthRadius

	x := 1
	fmt.Println(x)
	prt := &x
	*prt = 5

	fmt.Println(x)
	fmt.Println(prt)
}
