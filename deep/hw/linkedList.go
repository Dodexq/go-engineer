package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) listAppend(data int) {
	newNode := &Node{data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (list *LinkedList) listPrepend(data int) {
	if list.head == nil {
		newNode := &Node{data: data, next: nil}
		list.head = newNode
		return
	}
	newNode := &Node{data: data, next: list.head}
	list.head = newNode
}

func main() {
	a := &LinkedList{}
	a.listAppend(1)
	a.listAppend(2)
	a.listAppend(3)
	a.listPrepend(0)

	b := a.head
	for a.head.next != nil {
		fmt.Println(b)
		if b.next == nil {
			return
		}
		b = b.next

	}
}
