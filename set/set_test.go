package set

import (
    "fmt"
    "testing"
)

func Test(t *testing.T) {
    set := New()
    fmt.Printf("%v\n", set)

    set.Add(100)
    set.Add(200)
    fmt.Printf("%v\n", set.Contains(200))
    fmt.Printf("%v\n", set)
    set.Delete(200)
    fmt.Printf("%v\n", set)

    set2 := New()
    set2.Add(200)
    set2.Add(300)
    fmt.Printf("%v\n", set2)

    set.Union(set2)
    fmt.Printf("Union: %v\n", set)
    fmt.Printf("Set:%v\n", set)

    set = New()
    set.Add(1)
    set.Add(2)
    set.Add(3)

    set2 = New()
    set2.Add(1)
    set2.Add(4)
    set2.Add(5)
    set2.Add(2)

    fmt.Printf("Diff:%v\n", set.Diff(set2))
    fmt.Printf("Diff:%v\n", set2.Diff(set))

    fmt.Printf("Inter:%v\n", set.Inter(set2))

    fmt.Printf("SymDiff:%v\n", set.SymDiff(set2))

    set = New(0, 3, 6)
    fmt.Printf("Set:%v\n", set)
    set.Union(set2)
    fmt.Printf("Set:%v\n", set)

    set = New(0, 3, 6)
    set.Union2(set2)
    fmt.Printf("Set:%v\n", set)

}