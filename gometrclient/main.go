package main

import (
	"fmt"
	"gometrclient/gometr"
)

func main() {
	a := gometr.GoMetrClient{}
	a.NewGometr("https://ya.ru", 20)
	fmt.Println(a)

	ch := &gometr.Checker{}
	fz := &gometr.Checker{}
	ch.Add(ch)
	ch.Add(fz)
	fmt.Println(ch)
}
