package main

import (
	"fmt"
	"strings"
)

func addPrefix(origin string) string {
	return "prefix_" + origin
}

func addPrefixWithErr(origin string) (string, error) {
	return "prefix_" + origin, nil
}

func addPrefixWithLen(origin string) (res string, length int) {
	res = "prefix_" + origin
	length = len(origin)
	return
}

func main() {
	fmt.Println(addPrefix("string"))

	stringWithLen, le := addPrefixWithLen("Hello")
	fmt.Println(stringWithLen[7:], "have length", le, "bytes")

	s := strings.Map(func(r rune) rune { return r + 1 }, "SDWS")
	fmt.Println(s)

	var f, f2 func(s string) int
	f = func(s string) int { return len(s) }
	fmt.Println(f("string"))
	_ = f2

}
