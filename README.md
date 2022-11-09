# Circle Queue Go (CircleQueueGo)

The **Circle Queue** implementation in Go. The circle queue is a special version of queue where the last element of the queue is connected to the first element of the queue forming a circle. The operations are performed based on **FIFO** (First In First Out) principle. It is also called **'Ring Buffer'**.

## Installation

```bash
$ go get github.com/BaseMax/CircleQueueGo
```

## Usage

```go
package main

import "fmt"

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
```

© Copyright 2022, Max Base
