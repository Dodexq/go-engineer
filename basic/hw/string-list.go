package main

import "fmt"

func main() {
	var myList []string = []string{"a", "b", "c"}
	for count, status := range myList {
		_ = status
		if count+1 != len(myList) && myList[count] < myList[count+1] {
			fmt.Printf("Символ %s больше %s \n", myList[count], myList[count+1])
		}
	}
}
