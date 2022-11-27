package hashtable

import (
	"fmt"

	studentPkg "github.com/Chriszao/hashTable/student"
)

type Hash struct {
	maxItems     int
	maxPositions int
	size         int
	structure    []studentPkg.Student
}

func New(size int, max int) *Hash {
	structure := make([]studentPkg.Student, size)

	for index := 0; index < size; index++ {
		structure[index] = *studentPkg.New()
	}

	return &Hash{
		maxItems:     max,
		maxPositions: size,
		size:         0,
		structure:    structure,
	}
}

func (h *Hash) hashFunction(student studentPkg.Student) int {
	return student.GetRA() % h.maxPositions
}

func (h *Hash) DeleteHash() *Hash {
	h = nil

	return h
}

func (h *Hash) IsFull() bool {
	return h.size == h.maxItems
}

func (h *Hash) Len() int {
	return h.size
}

func (h *Hash) Insert(student studentPkg.Student) {
	index := h.hashFunction(student)

	h.structure[index] = student
	h.size++
}

func (h *Hash) Remove(student studentPkg.Student) {
	index := h.hashFunction(student)

	if h.structure[index].GetRA() != -1 {
		h.structure[index] = *studentPkg.NewWithParams(-1, " ")
		h.size--
	}
}

func (h *Hash) Find(student *studentPkg.Student, found *bool) {
	index := h.hashFunction(*student)
	aux := h.structure[index]

	if aux.GetRA() != student.GetRA() {
		*found = false
	} else {
		*found = true
		*student = aux
	}
}

func (h *Hash) Print() {
	printValue := "Hash table: \n"

	for index := 0; index < h.maxPositions; index++ {
		if h.structure[index].GetRA() != -1 {
			printValue += fmt.Sprintf(
				"Position: %d, RA: %d, Name: %s\n",
				index,
				h.structure[index].GetRA(),
				h.structure[index].GetName(),
			)
		}
	}

	fmt.Println(printValue)
}
