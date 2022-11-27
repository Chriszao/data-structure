package stack

import "fmt"

type Item int

const maxItems int = 100

type Stack struct {
	size      int
	structure []Item
}

// Constructs a new Stack
func New() *Stack {
	structure := make([]Item, maxItems)

	return &Stack{
		size:      0,
		structure: structure,
	}
}

// Delete the current Stack
func DeleteStack() *Stack {
	return nil
}

// Verify if stack is full
func (stack *Stack) IsFull() bool {
	return stack.size == maxItems
}

// Verify if stack is empty
func (stack *Stack) IsEmpty() bool {
	return stack.size == 0
}

// Add an item at the top of the stack
func (stack *Stack) Push(item Item) {
	if stack.IsFull() {
		fmt.Println("The stack is full!")
		return
	}

	stack.structure[stack.size] = item
	stack.size++
}

// Remove an item at the top of the stack. If the stack is empty, returns -1
func (stack *Stack) Pop() Item {
	if stack.IsEmpty() {
		fmt.Println("Stack is empty!")
		return -1
	}

	stack.size--
	removedElement := stack.structure[stack.size]
	stack.structure[stack.size] = 0

	return removedElement
}

// Prints the stack
func (stack *Stack) Print() {
	stackPrint := "[ "

	for index := 0; index < stack.size; index++ {
		stackPrint += fmt.Sprintf("%d ", stack.structure[index])
	}

	fmt.Println(stackPrint + "]")
}

// Returns the size of the stack
func (stack *Stack) Len() int {
	return stack.size
}
