package main

import (
	"fmt"
	"gometrclient/gometr"
)

func main() {
	a := gometr.GoMetrClient{}
	a.NewGometr("https://ya.ru", 20)
	fmt.Println(a)
	c := &gometr.Metrics{ID: "sd"}
	fmt.Println(c.GetID(), c.Health())
}
