package main

import "fmt"

// Main function
func main() {
	q := NewCircleQueue(5)
	fmt.Println(q)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	q.Enqueue(6)
	q.Enqueue(7)
	q.Enqueue(8)
	q.Enqueue(9)
	q.Enqueue(10)
	q.Enqueue(11)
	fmt.Println(q)
}
