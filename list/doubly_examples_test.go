package doublylinkedlist

import (
    "fmt"
)

func ExampleSlice() {
    list := New(1, 2, 3)
    fmt.Printf("%#v\n", list.Slice())
    // Output:
    // []interface {}{1, 2, 3}
}

func ExampleNodeString() {
    list := New("1")
    fmt.Println(list.First().String())
    // Output:
    // Node(data="1")
}

func ExampleListString() {
    list := New(1, 2, 3)
    list1 := New("1", "2", "3")
    list2 := New()
    fmt.Println(list.String())
    fmt.Println(list1.String())
    fmt.Println(list2.String())
    // Output:
    // [1, 2, 3]
    // ["1", "2", "3"]
    // []

}

func ExampleForeach() {
    list := New(1, 2, 3)
    list.ForEach()
    // Output:
    // 1
    // 2
    // 3

}

func ExampleForeachWithFunc() {
    list := New("1", "2", "3")
    list.ForEach(func(n ...interface{}) (int, error) {
        return fmt.Printf("%v\n", n...)
    })
    // Output:
    // Node(data="1")
    // Node(data="2")
    // Node(data="3")

}

func ExampleValue() {
    list := New(1, 2, 3)
    got := list.Front()
    fmt.Println(got.Value())
    // Output:
    // 1
}

func ExampleFirst() {
    list := New(1, 2, 3)
    got := list.First()
    fmt.Println(got.Value())
    // Output:
    // 1
}

func ExampleLast() {
    list := New(1, 2, 3)
    got := list.Last()
    fmt.Println(got.Value())
    // Output:
    // 3
}

func ExampleFind() {
    list := New(1, 2, 3)
    got := list.Find(2)
    fmt.Println(got)
    // Output:
    // Node(data=2)
}

func ExampleHas() {
    list := New(1, 2, 3)
    got := list.Has(2)
    fmt.Println(got)
    // Output:
    // true
}

func ExampleRemove() {
    list := New(1, 2, 3, 4)
    list.Remove(1)
    list.Remove(3)
    list.ForEach()
    // Output:
    // 2
    // 4
}

func ExampleRemove2() {

    list := New()
    list.Add(3, 3, 3)
    removed := list.Remove(3)
    fmt.Println(removed)
    list.ForEach()
    // Output:
    // true
    // 3
    // 3
}

func ExampleRemoveNode() {
    list := New(1, 2, 3, 4)
    list.RemoveNode(list.First())
    list.ForEach()
    // Output:
    // 2
    // 3
    // 4
}

func ExampleRemoveAll() {

    list := New()
    list.Add(3, 3, 3)
    list.RemoveAll(3)
    fmt.Println(list.String())
    // Output:
    // []
}

func ExampleRemoveAll2() {

    list := New()
    list.Add(1, 3, 3)
    list.RemoveAll(3)
    list.ForEach()
    // Output:
    // 1
}

func ExampleInsertBefore() {
    list := New(1, 2, 3)
    list.InsertBefore(0, list.First())
    list.ForEach()
    // Output:
    // 0
    // 1
    // 2
    // 3

}

func ExampleInsertAfter() {
    list := New(1, 2, 3)
    list.InsertAfter(4, list.Back())
    list.ForEach()
    // Output:
    // 1
    // 2
    // 3
    // 4

}

func ExampleMoveToFront() {
    list := New(1, 2, 3)
    node := list.Find(3)
    moved := list.MoveToFront(node)
    fmt.Println(moved)
    list.ForEach()
    // Output:
    // true
    // 3
    // 1
    // 2
}

func ExampleMoveToBack() {
    list := New(1, 2, 3)
    node := list.Find(1)
    moved := list.MoveToBack(node)
    fmt.Println(moved)
    list.ForEach()
    // Output:
    // true
    // 2
    // 3
    // 1
}

func ExampleMoveAfter() {
    list := New(1, 2, 3)
    node := list.Find(1)
    moved := list.MoveAfter(node, list.Back())
    fmt.Println(moved)
    list.ForEach()
    // Output:
    // true
    // 2
    // 3
    // 1
}

func ExampleMoveBefore() {
    list := New(1, 2, 3)
    node := list.Find(1)
    moved := list.MoveBefore(node, list.Back())
    fmt.Println(moved)
    list.ForEach()
    // Output:
    // true
    // 2
    // 1
    // 3
}

func ExampleReverse() {

    list := New()
    list.Add(1, 2, 3)
    list.Reverse()
    list.ForEach()
    // Output:
    // 3
    // 2
    // 1
}

func ExampleSwap() {

    list := New()
    list.Add(1, 2, 3)
    swapped := list.SwapValue(list.First(), list.Last())
    fmt.Println(swapped)
    list.ForEach()
    // Output:
    // true
    // 3
    // 2
    // 1
}

func ExamplePrepend() {

    list := New()
    list.Add(4, 5, 6)
    list.Prepend(1, 2, 3)
    list.ForEach()
    // Output:
    // 1
    // 2
    // 3
    // 4
    // 5
    // 6
}

func ExamplePushBackList() {
    list := New(1, 2, 3)
    list2 := New(4, 5, 6)
    list.PushBackList(list2)
    list.ForEach()
    // Output:
    // 1
    // 2
    // 3
    // 4
    // 5
    // 6
}

func ExamplePushFrontList() {
    list := New(4, 5, 6)
    list2 := New(1, 2, 3)
    list.PushFrontList(list2)
    list.ForEach()
    // Output:
    // 1
    // 2
    // 3
    // 4
    // 5
    // 6
}

func ExampleRemoveFront() {

    list := New(1, 2, 3)
    list.RemoveFront()
    fmt.Println(list.String())
    // Output:
    // [2, 3]
}

func ExampleRemoveBack() {

    list := New(1, 2, 3)
    list.RemoveBack()
    fmt.Println(list.String())
    // Output:
    // [1, 2]
}

func ExampleClear() {

    list := New(1, 2, 3)
    list.Clear()
    fmt.Println(list.String())
    // Output:
    // []
}
