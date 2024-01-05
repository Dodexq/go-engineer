package main

import "fmt"

type animal interface {
	walker
	runner
}

type bird interface {
	walker
	flier
}

type walker interface {
	walk()
}

type runner interface {
	run()
}

type flier interface {
	fly()
}

type cat struct{}
type eagle struct{}

func (c *cat) walk() {
	fmt.Println("Cat is walking")
}

func (c *cat) run() {
	fmt.Println("Cat is running")
}

func (e *eagle) walk() {
	fmt.Println("Eagle is walking")
}

func (c *eagle) fly() {
	fmt.Println("Eagle is flier")
}

func walk(w walker) {
	w.walk()
}

func fly(f flier) {
	f.fly()
}

func main() {
	var c animal = &cat{}
	var e bird = &eagle{}
	walk(c)
	walk(e)
	fly(e)
}
