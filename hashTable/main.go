package main

import (
	"fmt"
	"strconv"

	hashtable "github.com/Chriszao/hashTable/hashTable"
	"github.com/Chriszao/hashTable/student"
)

var raInput string
var nameInput string
var found bool = false

func main() {
	fmt.Println("**Hash generator**")

	var sizeInput, maxInput string

	fmt.Print("\nEnter the hash size: ")
	fmt.Scanln(&sizeInput)

	fmt.Print("\nEnter the max size of elements: ")
	fmt.Scanln(&maxInput)

	max, _ := strconv.Atoi(maxInput)
	size, _ := strconv.Atoi(sizeInput)

	fmt.Printf("\nThe load factor is: %.2f\n", float32(max)/float32(size))

	studentHash := hashtable.New(size, max)

	var option string

	for {

		fmt.Println("Enter 0 to stop the programm!")
		fmt.Println("\nEnter 1 to insert one student!")
		fmt.Println("\nEnter 2 to remove one student!")
		fmt.Println("\nEnter 3 to find one student!")
		fmt.Println("\nEnter 4 to print the hash!")

		fmt.Print("\nSelect: ")
		fmt.Scanln(&option)

		handleInput(option, studentHash)

		if option == "0" {
			fmt.Println("\nBye!")
			break
		}
	}
}

func handleInput(
	option string,
	studentHash *hashtable.Hash,
) {
	switch option {
	case "1":
		fmt.Print("\nWhat's the student RA: ")
		fmt.Scanln(&raInput)

		fmt.Print("\nWhat's the student Name: ")
		fmt.Scanln(&nameInput)

		ra, _ := strconv.Atoi(raInput)

		student := student.NewWithParams(ra, nameInput)

		studentHash.Insert(*student)
	case "2":
		fmt.Print("\nWhat's the student RA: ")
		fmt.Scanln(&raInput)

		ra, _ := strconv.Atoi(raInput)

		student := student.NewWithParams(ra, " ")

		studentHash.Remove(*student)

	case "3":
		fmt.Print("\nWhat's the student RA: ")
		fmt.Scanln(&raInput)

		ra, _ := strconv.Atoi(raInput)

		student := student.NewWithParams(ra, " ")

		studentHash.Find(student, &found)
		fmt.Println("Found", found)

		if found {
			fmt.Printf(
				"\nFound Student: \nRA: %d - Name: %s\n",
				student.GetRA(),
				student.GetName(),
			)
		} else {
			fmt.Printf("\nStudent not found!\n")
		}

	case "4":
		studentHash.Print()
	}
}
