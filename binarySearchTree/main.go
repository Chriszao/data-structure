package main

import (
	"fmt"
	"strconv"

	binarysearchtree "github.com/Chriszao/binarySearchTree/binarySearchTree"
	"github.com/Chriszao/binarySearchTree/student"
)

var raInput string
var nameInput string
var found bool = false

func main() {
	fmt.Println("**Binary Search Tree**")

	binarySearchTree := binarysearchtree.New()

	var option string

	for {

		fmt.Println("Enter 0 to stop the programm!")
		fmt.Println("\nEnter 1 to insert one student!")
		fmt.Println("\nEnter 2 to remove one student!")
		fmt.Println("\nEnter 3 to find one student!")
		fmt.Println("\nEnter 4 to print the tree!")

		fmt.Print("\nSelect: ")
		fmt.Scanln(&option)

		handleInput(option, binarySearchTree)

		if option == "0" {
			fmt.Println("\nBye!")
			break
		}
	}
}

func handleInput(
	option string,
	binarySearchTree *binarysearchtree.BinarySearchTree,
) {
	switch option {
	case "1":
		fmt.Print("\nWhat's the student Name: ")
		fmt.Scanln(&nameInput)

		fmt.Print("\nWhat's the student RA: ")
		fmt.Scanln(&raInput)

		ra, _ := strconv.Atoi(raInput)

		student := student.NewWithParams(ra, nameInput)

		binarySearchTree.Insert(*student)

	case "2":
		fmt.Print("\nWhat's the student RA: ")
		fmt.Scanln(&raInput)

		ra, _ := strconv.Atoi(raInput)

		student := student.NewWithParams(ra, "")

		binarySearchTree.Remove(*student)

	case "3":
		fmt.Print("\nWhat's the student RA: ")
		fmt.Scanln(&raInput)

		ra, _ := strconv.Atoi(raInput)

		student := student.NewWithParams(ra, "")

		binarySearchTree.Find(student, &found)

		if found {
			fmt.Println("***Found student***")
			fmt.Printf("Name: %s - RA: %d\n", student.GetName(), student.GetRa())
		} else {
			fmt.Println("Student no found!")
		}

	case "4":
		var impressionType string

		fmt.Println("Enter 1 to print the tree in pre-order;")
		fmt.Println("Enter 2 to print the tree in-order;")
		fmt.Println("Enter 3 to print the tree in post-order.")

		fmt.Print("input: ")
		fmt.Scanln(&impressionType)

		switch impressionType {
		case "1":
			binarySearchTree.PrintPreOrder(binarySearchTree.GetRoot())
		case "2":
			binarySearchTree.PrintInOrder(binarySearchTree.GetRoot())
		case "3":
			binarySearchTree.PrintPostOrder(binarySearchTree.GetRoot())
		}
	}
}
