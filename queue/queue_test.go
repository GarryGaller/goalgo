package queue

import (
    "fmt"
    "testing"
)


func Test(t *testing.T) {
    queue := New()
    fmt.Println(queue.Front())
    fmt.Println(queue.String(), queue.Size())
    queue.Enqueue(1)
    fmt.Println(queue.String(), queue.Size())
    queue.Enqueue(2)
    queue.Enqueue(3.5)
    queue.Enqueue(4)
    queue.Enqueue("���")
    queue.Enqueue("qwerty")
    fmt.Println(queue.Dequeue())

    fmt.Printf("%v => %d\n", queue.String(), queue.Size())
    //fmt.Println(queue.String())

}