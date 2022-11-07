/*
 * @Name: Circle Queue Go
 * @Author: Max Base
 * @Date: 2022-11-07
 * @Class: Data Structure, Dr. Mahsa Soheil Shamaee
 * @Repository: https://github.com/basemax/CircleQueueGo
 */
package main

import "fmt"

type CircleQueue struct {
	data  []interface{}
	rear  int
	front int
}

func NewCircleQueue(size int) *CircleQueue {
	return &CircleQueue{data: make([]interface{}, size), rear: 0, front: 0}
}

func (q *CircleQueue) IsFull() bool {
	return (q.rear+1)%len(q.data) == q.front
}

func (q *CircleQueue) IsEmpty() bool {
	return q.rear == q.front
}

func (q *CircleQueue) Enqueue(value interface{}) bool {
	if q.IsFull() {
		return false
	}
	q.data[q.rear] = value
	q.rear = (q.rear + 1) % len(q.data)
	return true
}

func (q *CircleQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	value := q.data[q.front]
	q.front = (q.front + 1) % len(q.data)
	return value
}

func (q *CircleQueue) Size() int {
	return (q.rear - q.front + len(q.data)) % len(q.data)
}

func (q *CircleQueue) Clear() {
	q.rear = 0
	q.front = 0
}

func (q *CircleQueue) String() string {
	result := ""
	for i := q.front; i != q.rear; i = (i + 1) % len(q.data) {
		result += fmt.Sprintf("%v ", q.data[i])
	}
	return result
}
