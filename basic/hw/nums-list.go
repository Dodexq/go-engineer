package main

import "fmt"

func main() {
	var nums []int = []int{1, 2, 56, 2, 56, 12, 52, 86, 3}
	var maps map[int]int = make(map[int]int)
	for _, y := range nums {
		maps[y]++
		if maps[y] >= 2 {
			fmt.Printf("Число %d встречается как минимум 2 раза\n", y)
		}
	}
}
