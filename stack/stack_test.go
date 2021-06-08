package queue

import (
    "fmt"
    "testing"
)



func Test(t *testing.T) {
    stack := New()
    fmt.Println(stack.Pop())
    fmt.Println(stack.String(), stack.Size())
    stack.Push(1)
    fmt.Println(stack.String(), stack.Size())
    stack.Push(2)
    stack.Push(3.5)
    stack.Push(4)
    stack.Push("сто")
    stack.Push("qwerty")
    fmt.Println(stack.Pop())

    fmt.Printf("%v => %d\n", stack.String(), stack.Size())
    //fmt.Println(stack.String())

}