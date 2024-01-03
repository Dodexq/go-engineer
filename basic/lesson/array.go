package main

import "fmt"

func main() {
	var x1 [5]int
	x1 = [5]int{1, 2, 3}
	fmt.Println(x1)

	var x2 [3]int = [3]int{1, 2, 3}
	fmt.Println(x2)

	x3 := [3]int{5, 6, 7}
	fmt.Println(x3)

	x4 := [...]float64{2.3, 2.2}
	fmt.Println(len(x4))

	x5 := "abv"
	fmt.Println(x5[1:])

	x6 := [...]string{"baba", "баба"}
	fmt.Println(len(x6[1]))
}
