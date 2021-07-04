## Linked-list, Queue and Stack implementation in Golang  




### Linked List
```go
---------------------
| 1 | 2 | 3 | 4 | 5 |
---------------------

Linked list is like an array;
an array is represented in memory like a chocolate bar
where each square piece is the value of the array;

but when we cut each piece and isolate it from the other we get a linked list

each node (piece) of the list links to the next one, and last node links to null
```  



### Queue
```go
Queue:
     ---------------------
Head | 1 | 2 | 3 | 4 | 5 | Tail
     ---------------------
FIFO: First In First Out

Queue is like a linked-list but instead of removing whatever you want
you can only remove from the head until you reach the end of the queue

inserting always from the tail of the queue
and deleting awalys from the head
```  

### Stack

```go

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
```

