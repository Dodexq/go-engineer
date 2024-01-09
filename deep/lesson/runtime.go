package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.Version())
	fmt.Println(runtime.NumCPU())
}
