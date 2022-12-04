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

func deleteTree(currentNode *Node) *BinarySearchTree {
	return &BinarySearchTree{
		root: nil,
	}
}

// Constructs a new BinarySearchTree
func New() *BinarySearchTree {
	return &BinarySearchTree{
		root: nil,
	}
}

func (bst *BinarySearchTree) Delete(currentNode *Node) *BinarySearchTree {
	return &BinarySearchTree{
		root: nil,
	}
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

func (bst *BinarySearchTree) Remove(student student.Student) {}

func (bst *BinarySearchTree) PrintPreOrder(currentNode *Node) {}

func (bst *BinarySearchTree) PrintInOrder(currentNode *Node) {}

func (bst *BinarySearchTree) PrintPostOrder(currentNode *Node) {}
