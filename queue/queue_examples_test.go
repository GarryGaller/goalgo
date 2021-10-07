package queue

import (
	"fmt"
)

func ExampleString() {

	queue := New()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue.String())
	// Output:
	// [1, 2, 3]

}

func ExampleEnqueue() {

	queue := New()
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
	}

	fmt.Println(queue)
	// Output:
	// [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

}

func ExampleDequeue() {

	queue := New()
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
	}

	for !queue.IsEmpty() {
		queue.Dequeue()
	}

	fmt.Println(queue)
	// Output:
	// []

}

func ExampleDequeue2() {

	queue := New()
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
	}

	for i := 0; i < 5; i++ {
		queue.Dequeue()
	}

	fmt.Println(queue)
	// Output:
	// [5, 6, 7, 8, 9]

}
