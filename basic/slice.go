package main

import "fmt"

func main() {
	var list []string
	fmt.Println(list)
	list = append(list, "string1", "5")
	fmt.Println(len(list[0]))

	var list1 []int64
	l := len(list1)
	c := cap(list1)

	fmt.Println(l)
	fmt.Println(c)

	list1 = []int64{1, 2, 3, 4, 5}
	l = len(list1)
	c = cap(list1)

	fmt.Println(l)
	fmt.Println(c)

	list1 = make([]int64, 2, 5)
	c = cap(list1)
	fmt.Println(list1)
	fmt.Println(c)

	list2 := make([]int64, 2, 3)
	fmt.Println(list2)
	fmt.Println(cap(list2))
	fmt.Println(len(list2))

	list2 = append(list2, 3)
	fmt.Println(list2)
	fmt.Println(cap(list2))

	list2 = append(list2, 4)
	fmt.Println(cap(list2))
}
