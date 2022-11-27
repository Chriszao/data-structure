package main

import (
	"fmt"
	"strconv"

	dinamicstackPkg "github.com/Chriszao/dinamicStack/dinamicStack"
)

func handleInput(
	input string,
	dinamicStack *dinamicstackPkg.DinamicStack,
) {
	switch input {
	case "1":
		var item string

		fmt.Print()
		fmt.Println("Enter the value to be inserted: ")
		fmt.Scanln(&item)
		stackItem, _ := strconv.Atoi(item)

		dinamicStack.Push(dinamicstackPkg.Item(stackItem))
	case "2":
		fmt.Println("Removed Item: ", dinamicStack.Pop())
	case "3":
		dinamicStack.Print()
	case "4":
		fmt.Println("Stack Size: ", dinamicStack.Len())
	case "5":
		fmt.Println("Stack is empty?: ", dinamicStack.IsEmpty())
	case "6":
		fmt.Println("Head of the stack: ", dinamicStack.Peek())
	}
}

func main() {
	dinamicStack := dinamicstackPkg.New()

	var input string

	for {
		fmt.Println("Enter 0 to stop the programm!")
		fmt.Println("\nEnter 1 to push one element!")
		fmt.Println("\nEnter 2 to pop one element!")
		fmt.Println("\nEnter 3 to print the stack!")
		fmt.Println("\nEnter 4 to print the stack size!")
		fmt.Println("\nEnter 5 to check if the stack is empty!")
		fmt.Println("\nEnter 6 to peek the stack!")

		fmt.Println()
		fmt.Print("input: ")
		fmt.Scanln(&input)
		handleInput(input, dinamicStack)

		if input == "0" {
			fmt.Println("\nBye!")
			break
		}
	}
}
