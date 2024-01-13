package mypackage

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) InsertAtBack(Data int) {
	newNode := &Node{Data: Data, Next: nil}
	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (list *LinkedList) InsertAtPrev(Data int) {
	if list.head == nil {
		list.head = &Node{Data: Data, Next: nil}
		return
	} else {
		newNode := &Node{Data: Data, Next: list.head}
		list.head = newNode
	}
}

func (list *LinkedList) ShowAll() {
	head := list.head
	if head.Next != nil {
		for head.Next != nil {
			fmt.Print(head.Data, " ")
			head = head.Next
		}
	}
	fmt.Println(head.Data)

}
