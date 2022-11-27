package queue

import "fmt"

type Item int

const maxItems int = 100

type Queue struct {
	first     int
	last      int
	structure []Item
}

// Contructs a new Queue
func New() *Queue {
	structure := make([]Item, maxItems)

	return &Queue{
		first:     0,
		last:      0,
		structure: structure,
	}
}

// Delete the current Queue
func (queue *Queue) DeleteQueue() *Queue {
	queue = nil

	return queue
}

// Check if queue is empty
func (queue *Queue) IsEmpty() bool {
	return queue.first == queue.last
}

// Check if queue is full
func (queue *Queue) IsFull() bool {
	return (queue.last - queue.first) == maxItems
}

// Insert one item at the end of the queue
func (queue *Queue) Enqueue(item Item) {
	if queue.IsFull() {
		fmt.Println("Queue is full!")
		return
	}

	currentIndex := queue.last % maxItems

	queue.structure[currentIndex] = item
	queue.last++
}

// Remove one item at the start of the queue, and return it
func (queue *Queue) Dequeue() Item {
	if queue.IsEmpty() {
		fmt.Println("Queue is empty, nothing removed!")
		return 0
	}

	currentindex := queue.first % maxItems
	queue.first++

	return queue.structure[currentindex]
}

// Return the size of the queue
func (queue *Queue) Len() int {
	return queue.last - queue.first
}

// Prints the queue
func (queue *Queue) Print() {
	printValue := "["

	for index := queue.first; index < queue.last; index++ {
		printValue += fmt.Sprintf(" %d ", queue.structure[index%maxItems])
	}

	fmt.Println(printValue + "]")
}
