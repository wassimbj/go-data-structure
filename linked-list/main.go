package main

/*
	Linked List:
	---------------------
	| 1 | 2 | 3 | 4 | 5 |
	---------------------

	Linked list is like an array;
	an array is represented in memory like a chocolate bar
	where each square piece is the value of the array;

	but when we cut each piece and isolate it from the other we get a linked list

	each node (piece) of the list points to the next one, and last node points to null
*/

import "fmt"

type Node struct {
	data string
	next *Node
}

type List struct {
	head *Node
}

func (l *List) isEmpty() bool {
	if l.head == nil {
		return true
	}

	return false
}

func (l *List) add(data string) {
	newNode := new(Node)

	newNode.data = data
	newNode.next = nil

	if l.isEmpty() {
		l.head = newNode
		return
	}
	// else
	lastNode := l.head
	for lastNode.next != nil {
		// next item, ....
		lastNode = lastNode.next
	}
	lastNode.next = newNode
}

func (l *List) del(data string) {
	nodeToDel := l.head
	prevNode := l.head

	for nodeToDel != nil {
		if nodeToDel.data == data {
			break
		} else {
			prevNode = nodeToDel
		}

		nodeToDel = nodeToDel.next
	}
	prevNode.next = nodeToDel.next
	nodeToDel = nil
}

func (l *List) print() {
	if l.isEmpty() {
		fmt.Print("List is empty, add items first")
		return
	}

	currentNode := l.head
	fmt.Print("------ Data Stored ---- \n")
	for currentNode != nil {
		fmt.Printf("%s \n", currentNode.data)
		currentNode = currentNode.next
	}
}

func main() {

	var list List
	list.head = nil

	list.add("user-1")
	list.add("user-2")
	list.add("user-3")
	list.add("user-200")
	list.add("user-300")

	list.print()
	list.del("user-3")
	list.print()

}
