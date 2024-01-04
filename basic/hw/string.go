package main

import "fmt"

func main() {
	myString := "съешь ещё этих мягких французских булок, да выпей чаю"

	myMap := make(map[string]int)
	for _, y := range myString {
		key := fmt.Sprintf("%c", y)
		if key != " " {
			myMap[key] += 1
		}
	}

	for i, y := range myMap {
		fmt.Println(i, "-", y)
	}
}
