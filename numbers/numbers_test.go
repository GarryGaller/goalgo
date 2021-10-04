package numbers

import (
    "fmt"
    "sort"
    "testing"
)

func TestNum(t *testing.T) {

    a := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
    n := Deduplicate(a)
    fmt.Printf("%v\n", a[:n])

    a = []int{2, 3, 4, 5, 7, 7, 2, 2, 3, 100}
    sort.Ints(a)
    n = Deduplicate2(a)
    fmt.Printf("%v\n", a[:n])

}
