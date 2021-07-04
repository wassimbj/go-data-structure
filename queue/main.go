package main

/*
	Queue:
	  ---------------------
Head | 1 | 2 | 3 | 4 | 5 | Tail
	  ---------------------
	FIFO: First In First Out

	Queue is like a linked-list but instead of removing whatever you want
	you can only remove from the head until you reach the end of the queue

	inserting always from the tail of the queue
	and deleting awalys from the head
*/

import "fmt"

type Node struct {
	data string
	next *Node
}

type Queue struct {
	head *Node
}

func (q *Queue) isEmpty() bool {
	if q.head == nil {
		return true
	}

	return false
}

func (q *Queue) add(data string) {
	newNode := new(Node)

	newNode.data = data
	newNode.next = nil

	if q.isEmpty() {
		q.head = newNode
		return
	}
	// else
	lastNode := q.head
	for lastNode.next != nil {
		// next item, ....
		lastNode = lastNode.next
	}
	lastNode.next = newNode
}

func (q *Queue) del() string {
	head := q.head

	if q.isEmpty() {
		fmt.Print("Queue is empty !!")
		return ""
	}

	deleted := q.head.data
	q.head = head.next
	head = nil

	return deleted
}

func (q *Queue) print() {
	if q.isEmpty() {
		fmt.Print("Queue is empty, add items first")
		return
	}

	currentNode := q.head
	fmt.Print("------ Data Stored ---- \n")
	for currentNode != nil {
		fmt.Printf("%s \n", currentNode.data)
		currentNode = currentNode.next
	}
}

func main() {

	var queue Queue
	queue.head = nil

	queue.add("user-1")
	queue.add("user-2")
	queue.add("user-3")

	queue.print()
	queue.del()
	queue.print()

}
