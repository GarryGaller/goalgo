//https://ewencp.org/blog/golang-iterators/index.html

package main

type StatefulIterator interface {
    Value() int
    Next() bool
}

type intStatefulIterator struct {
    current int
    data    []int
}

func (it *intStatefulIterator) Value() int {
    return it.data[it.current]
}
func (it *intStatefulIterator) Next() bool {
    it.current++
    if it.current >= len(it.data) {
        return false
    }
    return true
}
func NewIntStatefulIterator(data []int) *intStatefulIterator {
    return &intStatefulIterator{data: data, current: -1}
}


var sum int = 0
for it := NewIntStatefulIterator(int_data); it.Next(); {
    sum += it.Value()
}