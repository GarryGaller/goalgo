package numbers

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	a := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	n := removeDuplicates(a)
	fmt.Printf("%v\n", a[:n])

	a = []int{2, 3, 4, 5, 7, 7, 2, 2, 3, 100}
	sort.Ints(a)
	n = removeDuplicates(a)
	fmt.Printf("%v\n", a[:n])

	a = []int{2, 3, 4, 5, 7, 7, 2, 2, 3, 100}
	sort.Ints(a)
	n = deduplicate(a)
	fmt.Printf("%v\n", a[:n])

	a = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	sort.Ints(a)
	n = deduplicate(a)
	fmt.Printf("%v\n", a[:n])

}
