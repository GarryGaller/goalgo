package numbers_test

import (
    "fmt"
    "sort"
    "testing"
    "github.com/GarryGaller/goalgo/numbers"
)

func Test(t *testing.T) {

    a := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
    n := numbers.RemoveDuplicates(a)
    fmt.Printf("%v\n", a[:n])

    a = []int{2, 3, 4, 5, 7, 7, 2, 2, 3, 100}
    sort.Ints(a)
    n = numbers.RemoveDuplicates(a)
    fmt.Printf("%v\n", a[:n])

    a = []int{2, 3, 4, 5, 7, 7, 2, 2, 3, 100}
    sort.Ints(a)
    n = numbers.Deduplicate(a)
    fmt.Printf("%v\n", a[:n])

    a = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
    sort.Ints(a)
    n = numbers.Deduplicate(a)
    fmt.Printf("%v\n", a[:n])

}
