package main

import "fmt"

type greeter interface {
	greet(string) string
}

type russian struct{}
type american struct{}

func (r *russian) greet(name string) string {
	return fmt.Sprintln("Привет", name)
}

func (r *american) greet(name string) string {
	return fmt.Sprintln("Hello", name)
}

func sayHello(g greeter, name string) {
	fmt.Print(g.greet(name))
}

func main() {
	sayHello(&russian{}, "Артем")
	sayHello(&american{}, "Artem")
}
