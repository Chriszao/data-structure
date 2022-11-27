package main

import (
	"fmt"
	"strconv"

	"github.com/Chriszao/stack/stack"
)

func handleInput(
	input string,
	stackObj *stack.Stack,
) {
	switch input {
	case "1":
		var item string

		fmt.Print()
		fmt.Println("Enter the value to be inserted: ")
		fmt.Scanln(&item)
		stackItem, _ := strconv.Atoi(item)

		stackObj.Push(stack.Item(stackItem))
	case "2":
		fmt.Println("Removed Item: ", stackObj.Pop())
	case "3":
		stackObj.Print()
	case "4":
		fmt.Println("Stack Size: ", stackObj.Len())
	case "5":
		fmt.Println("Stack is empty?: ", stackObj.IsEmpty())
	case "6":
		fmt.Println("Stack is full?: ", stackObj.IsFull())
	}
}

func main() {
	stackObj := stack.New()

	var input string

	for {
		fmt.Println("Enter 0 to stop the programm!")
		fmt.Println("\nEnter 1 to push one element!")
		fmt.Println("\nEnter 2 to pop one element!")
		fmt.Println("\nEnter 3 to print the stack!")
		fmt.Println("\nEnter 4 to print the stack size!")
		fmt.Println("\nEnter 5 to check if the stack is empty!")
		fmt.Println("\nEnter 6 to check if the stack is full!")

		fmt.Println()
		fmt.Print("input: ")
		fmt.Scanln(&input)
		handleInput(input, stackObj)

		if input == "0" {
			fmt.Println("\nBye!")
			break
		}
	}
}
