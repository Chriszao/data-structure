package dinamicqueue

import "fmt"

type Item int

type Node struct {
	value Item
	next  *Node
}

type DinamicQueue struct {
	size  int
	rear  *Node
	front *Node
}

func New() *DinamicQueue {
	return &DinamicQueue{
		size:  0,
		rear:  nil,
		front: nil,
	}
}

func (dq *DinamicQueue) DeleteQueue() *DinamicQueue {
	for dq.front != nil {
		dq.front = dq.front.next
	}

	dq.rear = nil
	dq.size = 0

	return dq
}

func (dq *DinamicQueue) IsEmpty() bool {
	return dq.size == 0
}

func (dq *DinamicQueue) Enqueue(item Item) {
	newNode := &Node{
		value: item,
		next:  nil,
	}

	if dq.front == nil {
		dq.front = newNode
	} else {
		dq.rear.next = newNode
	}

	dq.rear = newNode
	dq.size++
}

func (dq *DinamicQueue) Dequeue() Item {
	if dq.IsEmpty() {
		fmt.Println("The Queue is empty!")
		return 0
	}

	item := dq.front.value
	dq.front = dq.front.next
	if dq.front == nil {
		dq.rear = nil
	}
	dq.size--

	return item
}

func (dq *DinamicQueue) Print() {
	temporaryNode := dq.front
	printValue := "Queue: ["

	for temporaryNode != nil {
		printValue += fmt.Sprintf(" %d ", temporaryNode.value)
		temporaryNode = temporaryNode.next
	}

	fmt.Println(printValue + "]")
}

func (dq *DinamicQueue) Peek() Item {
	return dq.front.value
}

func (dq *DinamicQueue) Len() int {
	return dq.size
}
