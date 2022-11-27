package dinamicstack

import "fmt"

type Item int

type Node struct {
	value Item
	next  *Node
}

type DinamicStack struct {
	head *Node
	size int
}

// Constructs a DinamicStack
func New() *DinamicStack {
	return &DinamicStack{
		head: nil,
		size: 0,
	}
}

// Delete the current stack
func (ds *DinamicStack) DeleteDinamicStack() *DinamicStack {
	for ds.head != nil {
		ds.head = ds.head.next
		ds.size--
	}

	return ds
}

// Return if the stack is empty
func (ds *DinamicStack) IsEmpty() bool {
	return ds.size == 0
}

// Peek the value of the head of the stack
func (ds *DinamicStack) Peek() Item {
	if ds.IsEmpty() {
		fmt.Println("Stack is empty!")
		return 0
	}

	return ds.head.value
}

// Insert a new value in the head of the stack
func (ds *DinamicStack) Push(item Item) {
	newNode := &Node{
		value: item,
		next:  ds.head,
	}

	ds.head = newNode
	ds.size++
}

// Remove the value at the head of the stack
func (ds *DinamicStack) Pop() Item {
	if ds.IsEmpty() {
		fmt.Println("The stack is empty!")
		return 0
	}

	item := ds.head.value
	ds.head = ds.head.next
	ds.size--

	return item
}

// Returns the size of the stack
func (ds *DinamicStack) Len() int {
	return ds.size
}

// Prints the stack
func (ds *DinamicStack) Print() {
	temporaryNode := ds.head
	printValue := "Stack: ["

	for temporaryNode != nil {
		printValue += fmt.Sprintf(" %d ", temporaryNode.value)

		temporaryNode = temporaryNode.next
	}

	fmt.Println(printValue + "]")
}
