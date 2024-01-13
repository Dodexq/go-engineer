package main

import (
	"fmt"
	"gometrclient/gometr"
)

func main() {
	c := gometr.Data{}
	c.IDs = c.GetMetrics()
	fmt.Print(c.IDs)
}
