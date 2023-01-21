package binarysearchtree

import (
	"fmt"

	"github.com/Chriszao/binarySearchTree/student"
)

type Node struct {
	student    student.Student
	leftChild  *Node
	rightChild *Node
}

type BinarySearchTree struct {
	root *Node
}

func (bst *BinarySearchTree) getSucessor(
	sucessorStudent *student.Student,
	temporaryNode *Node,
) {
	temporaryNode = temporaryNode.rightChild

	for temporaryNode.leftChild != nil {
		temporaryNode = temporaryNode.leftChild
	}

	*sucessorStudent = temporaryNode.student
}

func (bst *BinarySearchTree) deleteNode(currentNode *Node) {
	// Case 1: Has none or one child
	if currentNode.leftChild == nil {
		*currentNode = *currentNode.rightChild
		return
	}

	// Case 2: Case 1 and has only one child in right
	if currentNode.rightChild == nil {
		*currentNode = *currentNode.leftChild
		return
	}

	// Case 3: Has two childrens
	var sucessorStudent student.Student

	bst.getSucessor(&sucessorStudent, currentNode)

	currentNode.student = sucessorStudent

	bst.removeSearch(sucessorStudent, currentNode.rightChild)
}

func (bst *BinarySearchTree) removeSearch(
	student student.Student,
	currentNode *Node,
) {
	if student.GetRa() < currentNode.student.GetRa() {
		bst.removeSearch(student, currentNode.leftChild)
		return
	}

	if student.GetRa() > currentNode.student.GetRa() {
		bst.removeSearch(student, currentNode.rightChild)
		return
	}

	bst.deleteNode(currentNode)
}

func deleteTree(currentNode *Node) {
	if currentNode == nil {
		deleteTree(currentNode.leftChild)

		deleteTree(currentNode.rightChild)

		currentNode = nil
	}
}

// Constructs a new BinarySearchTree
func New() *BinarySearchTree {
	return &BinarySearchTree{
		root: nil,
	}
}

func (bst *BinarySearchTree) DeleteTree(currentNode *Node) {
	deleteTree(bst.root)
}

func (bst *BinarySearchTree) GetRoot() *Node {
	return bst.root
}

func (bst *BinarySearchTree) IsEmpty() bool {
	return bst.root == nil
}

func (bst *BinarySearchTree) Insert(student student.Student) {
	newNode := &Node{
		student:    student,
		leftChild:  nil,
		rightChild: nil,
	}

	if bst.root == nil {
		bst.root = newNode
		return
	}

	temporaryNode := bst.root

	for temporaryNode != nil {
		if student.GetRa() < temporaryNode.student.GetRa() {
			if temporaryNode.leftChild == nil {
				temporaryNode.leftChild = newNode
				break
			}

			temporaryNode = temporaryNode.leftChild
		} else {
			if temporaryNode.rightChild == nil {
				temporaryNode.rightChild = newNode
				break
			}

			temporaryNode = temporaryNode.rightChild
		}
	}
}

func (bst *BinarySearchTree) Find(student *student.Student, found *bool) {
	if bst.IsEmpty() {
		*found = false
		fmt.Println("Binary Search Tree is empty")
		return
	}

	currentNode := bst.root

	for currentNode != nil {
		if student.GetRa() < currentNode.student.GetRa() {
			currentNode = currentNode.leftChild
			continue
		}

		if student.GetRa() > currentNode.student.GetRa() {
			currentNode = currentNode.rightChild
			continue
		}

		*found = true
		*student = currentNode.student
		break
	}
}

func (bst *BinarySearchTree) Remove(student student.Student) {
	if bst.IsEmpty() {
		fmt.Println("Binary Search Tree is empty")
		return
	}

	bst.removeSearch(student, bst.root)
}

func (bst *BinarySearchTree) PrintPreOrder(currentNode *Node) {

	if currentNode != nil {
		fmt.Printf("%s: %d\n",
			currentNode.student.GetName(),
			currentNode.student.GetRa(),
		)

		bst.PrintPreOrder(currentNode.leftChild)

		bst.PrintPreOrder(currentNode.rightChild)
	}
}

func (bst *BinarySearchTree) PrintInOrder(currentNode *Node) {

	if currentNode != nil {
		bst.PrintInOrder(currentNode.leftChild)

		fmt.Printf("%s: %d\n",
			currentNode.student.GetName(),
			currentNode.student.GetRa(),
		)

		bst.PrintInOrder(currentNode.rightChild)
	}
}

func (bst *BinarySearchTree) PrintPostOrder(currentNode *Node) {

	if currentNode != nil {
		bst.PrintPostOrder(currentNode.leftChild)

		bst.PrintPostOrder(currentNode.rightChild)

		fmt.Printf("%s: %d\n",
			currentNode.student.GetName(),
			currentNode.student.GetRa(),
		)
	}
}
