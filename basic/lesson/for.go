package main

import "fmt"

func main() {
	mySlice := []string{"A", "B", "C"}
	fmt.Println(mySlice)

	for key, value := range mySlice {
		fmt.Printf("key %v, val: %v \n", key, value)
	}

	for _, value := range mySlice {
		fmt.Println(value)
	}

	ages := map[string]int{
		"Алексей": 24,
		"Виктор":  43,
	}
	fmt.Println(ages)

	for name := range ages {
		fmt.Println(name)
	}

	someString := "дваБайта"
	for i := 0; i < len(someString); i++ {
		fmt.Print(someString[i], " ")
		fmt.Printf("%T ", someString[i])
	}

	x := 42
	fmt.Printf("%T\n", x)

	for key, value := range someString {
		fmt.Println(key)
		fmt.Println(value)
		fmt.Printf("%T", value)
	}
}
