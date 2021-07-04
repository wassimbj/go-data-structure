package main

/*
	Stack:

	 Head
	|---|
	| 4 |
	|---|
	| 3 |
	|---|
	| 2 |
	|---|
	| 1 |
	|---|
	 Tail

	LIFO: Last In First Out

	Stack is like a Queue but reversed

	inserting always from the head
	and deleting also from the head

	its like a stack of books or coins
*/

import "fmt"

type Node struct {
	data string
	next *Node
}

type Stack struct {
	head *Node
}

func (s *Stack) isEmpty() bool {
	if s.head == nil {
		return true
	}

	return false
}

func (s *Stack) add(data string) {
	newNode := new(Node)

	newNode.data = data
	newNode.next = s.head
	s.head = newNode
}

func (s *Stack) del() string {
	head := s.head

	if s.isEmpty() {
		fmt.Print("Stack is empty !!")
		return ""
	}

	deleted := s.head.data
	s.head = head.next
	head = nil

	return deleted
}

func (s *Stack) print() {
	if s.isEmpty() {
		fmt.Print("Stack is empty, add items first")
		return
	}

	currentNode := s.head
	fmt.Print("------ Data Stored ---- \n")
	for currentNode != nil {
		fmt.Printf("%s \n", currentNode.data)
		currentNode = currentNode.next
	}
}

func main() {

	var stack Stack
	stack.head = nil

	stack.add("user-1")
	stack.add("user-2")
	stack.add("user-3")

	stack.print()
	// stack.del()
	// stack.print()

}
