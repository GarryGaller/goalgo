package sort

import (
    "testing"
    "reflect"
)

func TestCyclicSort(t *testing.T) {

    var testCases = []struct {
        in       []int
        expected []int
    }{
        {
            []int{6, 4, 5, 2, 3, 1},
            []int{1, 2, 3, 4, 5, 6},
        },

        {
            []int{4, 5, 2, 3, 1},
            []int{1, 2, 3, 4, 5},
        },

        {
            []int{2, 3, 1},
            []int{1, 2, 3},
        },
    }

    for _, c := range testCases {
        СyclicSort(c.in)
        if !reflect.DeepEqual(c.in,c.expected) {
            t.Errorf("CyclicSort() = %v; want %v", c.in, c.expected)
        }
    }
}

func TestСountingSort(t *testing.T) {

    var testCases = []struct {
        in       []int
        expected []int
    }{
        {
            []int{3, 1, 2},
            []int{1, 2, 3},
        },

        {
            []int{10, 5, 0},
            []int{0, 5, 10},
        },

        {
            []int{10, 9, 8},
            []int{8, 9, 10},
        },
    }

    for _, c := range testCases {
        got := СountingSort(c.in, 10)
        if !reflect.DeepEqual(got, c.expected) {
            t.Errorf("CountingSort() = %v; want %v", got, c.expected)
        }
    }
}   


func TestSelectionSort(t *testing.T) {

    var testCases = []struct {
        in       []int
        expected []int
    }{
        {
            []int{3, 1, 2},
            []int{1, 2, 3},
        },

        {
            []int{10, 5, 0},
            []int{0, 5, 10},
        },

        {
            []int{10, 9, 8},
            []int{8, 9, 10},
        },
    }

    for _, c := range testCases {
        
        SelectionSort(c.in)
        if !reflect.DeepEqual(c.in, c.expected) {
            t.Errorf("SelectionSort() = %v; want %v", c.in, c.expected)
        }
    }
}


func TestInsertionSort(t *testing.T) {

    var testCases = []struct {
        in       []int
        expected []int
    }{
        {
            []int{3, 1, 2},
            []int{1, 2, 3},
        },

        {
            []int{10, 5, 0},
            []int{0, 5, 10},
        },

        {
            []int{10, 9, 8},
            []int{8, 9, 10},
        },
    }

    for _, c := range testCases {
        
        InsertionSort(c.in)
        if !reflect.DeepEqual(c.in, c.expected) {
            t.Errorf("InsertionSort() = %v; want %v", c.in, c.expected)
        }
    }
}
  

func TestQuickSort(t *testing.T) {

    var testCases = []struct {
        in       []int
        expected []int
    }{
        {
            []int{3, 1, 2},
            []int{1, 2, 3},
        },

        {
            []int{10, 5, 0},
            []int{0, 5, 10},
        },

        {
            []int{10, 9, 8},
            []int{8, 9, 10},
        },
    }

    for _, c := range testCases {
        
        got := QuickSort(c.in)
        if !reflect.DeepEqual(got, c.expected) {
            t.Errorf("QuickSort() = %v; want %v", got, c.expected)
        }
    }
}  



func TestMergeSort(t *testing.T) {

    var testCases = []struct {
        in       []int
        expected []int
    }{
        {
            []int{3, 1, 2},
            []int{1, 2, 3},
        },

        {
            []int{10, 5, 0},
            []int{0, 5, 10},
        },

        {
            []int{10, 9, 8},
            []int{8, 9, 10},
        },
    }

    for _, c := range testCases {
        
        got := MergeSort(c.in)
        if !reflect.DeepEqual(got, c.expected) {
            t.Errorf("MergeSort() = %v; want %v", got, c.expected)
        }
    }
}
