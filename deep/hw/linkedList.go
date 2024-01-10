package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) insertAtBack(data int) {
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

func (list *LinkedList) insertAtPrev(data int) {
	if list.head == nil {
		list.head = &Node{data: data, next: nil}
		return
	} else {
		newNode := &Node{data: data, next: list.head}
		list.head = newNode
	}
}

func (list *LinkedList) showAll() {
	head := list.head
	if head.next != nil {
		for head.next != nil {
			fmt.Println(head.data)
			head = head.next
		}
	}
	fmt.Println(head.data)

}

func main() {
	l := &LinkedList{}
	l.insertAtBack(1)
	l.insertAtBack(2)
	l.insertAtPrev(5)
	l.insertAtBack(3)
	l.showAll()
}
