package main

import (
	"fmt"
	"strconv"

	dinamicqueuePkg "github.com/Chriszao/dinamicQueue/dinamicQueue"
)

func handleInput(
	input string,
	dinamicQueue *dinamicqueuePkg.DinamicQueue,
) {
	switch input {
	case "1":
		var item string

		fmt.Print()
		fmt.Println("Enter the value to be inserted: ")
		fmt.Scanln(&item)
		queueItem, _ := strconv.Atoi(item)

		dinamicQueue.Enqueue(dinamicqueuePkg.Item(queueItem))
	case "2":
		fmt.Println("Removed Item: ", dinamicQueue.Dequeue())
	case "3":
		dinamicQueue.Print()
	case "4":
		fmt.Println("queue Size: ", dinamicQueue.Len())
	case "5":
		fmt.Println("queue is empty?: ", dinamicQueue.IsEmpty())
	case "6":
		fmt.Println("Head of the queue: ", dinamicQueue.Peek())
	}
}

func main() {
	dinamicQueue := dinamicqueuePkg.New()

	var input string

	for {
		fmt.Println("Enter 0 to stop the programm!")
		fmt.Println("\nEnter 1 to enqueue one element!")
		fmt.Println("\nEnter 2 to dequeue one element!")
		fmt.Println("\nEnter 3 to print the queue!")
		fmt.Println("\nEnter 4 to print the queue size!")
		fmt.Println("\nEnter 5 to check if the queue is empty!")
		fmt.Println("\nEnter 6 to peek the queue!")

		fmt.Println()
		fmt.Print("input: ")
		fmt.Scanln(&input)
		handleInput(input, dinamicQueue)

		if input == "0" {
			fmt.Println("\nBye!")
			break
		}
	}
}
