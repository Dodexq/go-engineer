package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myString string = "\tHello\n"
	fmt.Println(myString)

	var testString byte = 1
	fmt.Println(reflect.TypeOf(testString))

	str := "12345"
	fmt.Println(str[2])

	str = "строка"
	fmt.Println(len(str))
	fmt.Println(str[0:6])
}
