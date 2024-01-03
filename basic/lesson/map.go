package main

import "fmt"

func main() {
	var m1 map[string]string
	fmt.Println(m1)

	m3 := make(map[int]int)
	m3[1] = 1
	fmt.Println(m3)

	ages := map[string]int{
		"Андрей": 20,
		"Артем":  28,
	}
	fmt.Println(ages)

	fmt.Println(ages["Артем"])
	ages["Андрей"] = 32
	fmt.Println(ages)

	ages["Михаил"] = 42
	fmt.Println(ages)

	ages["Владимир"]++
	fmt.Println(ages["Владимир"])
}
