package main

import (
	"fmt"
	"strconv"

	"github.com/Chriszao/queue/queue"
)

func handleInput(
	input string,
	queueObj *queue.Queue,
) {
	switch input {
	case "1":
		var item string

		fmt.Print()
		fmt.Println("Enter the value to be inserted: ")
		fmt.Scanln(&item)
		queueItem, _ := strconv.Atoi(item)

		queueObj.Enqueue(queue.Item(queueItem))
	case "2":
		fmt.Println("Removed Item: ", queueObj.Dequeue())
	case "3":
		queueObj.Print()
	case "4":
		fmt.Println("queue is empty?: ", queueObj.IsEmpty())
	case "5":
		fmt.Println("queue is full?: ", queueObj.IsFull())
	case "6":
		fmt.Println("Queue size: ", queueObj.Len())
	}
}

func main() {
	queueObj := queue.New()

	var input string

	for {
		fmt.Println("Enter 0 to stop the programm!")
		fmt.Println("\nEnter 1 to enqueue one element!")
		fmt.Println("\nEnter 2 to dequeue one element!")
		fmt.Println("\nEnter 3 to print the queue!")
		fmt.Println("\nEnter 4 to check if the queue is empty!")
		fmt.Println("\nEnter 5 to check if the queue is full!")
		fmt.Println("\nEnter 6 to check the queue size!")

		fmt.Println()
		fmt.Print("input: ")
		fmt.Scanln(&input)
		handleInput(input, queueObj)

		if input == "0" {
			fmt.Println("\nBye!")
			break
		}
	}
}
