package doublylinkedlist

//https://github.com/omerkaya1/doublylinkedlist/blob/master/list.go

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "fmt"
)

var (
    testData1 = "first"
    testData2 = "second"
    testData3 = "third"
    testData4 = 1
    testData5 = 2
    testData6 = 3
)

func cleanup(impl *List) {
    impl = nil
    return
}

/**
 * A set o general tests
 */
func TestNewList(t *testing.T) {
    list := New()
    defer cleanup(list)

    if assert.NotNil(t, list) {
        assert.Equal(t, 0, list.Len())
        assert.Nil(t, list.First())
        assert.Nil(t, list.Last())
    }
}

func TestListGeneral(t *testing.T) {
    list := New()
    defer cleanup(list)

    list.PushFront(testData1)
    list.PushFront(testData2)
    list.PushFront(testData3)

    list.PushBack(testData4)
    list.PushBack(testData5)
    list.PushBack(testData6)

    if assert.NotNil(t, list) {
        assert.Equal(t, 6, list.Len())
        assert.Equal(t, testData3, list.First().Data())
        assert.Equal(t, testData6, list.Last().Data())
    }
}

/**
 * A set of test for the list functionality
 */
func TestListImpl_SingleElementList(t *testing.T) {
    list := New()
    defer cleanup(list)

    list.PushFront(testData1)

    if assert.NotNil(t, list) {
        assert.Equal(t, 1, list.Len())
        assert.True(t, list.First() == list.Last())
    }
}

func TestListImpl_RemoveSingleElement(t *testing.T) {
    list := New()
    defer cleanup(list)

    list.PushFront(testData1)

    if assert.NotNil(t, list) {
        assert.Equal(t, 1, list.Len())
        assert.True(t, list.First() == list.Last())
        list.RemoveNode(list.First())
        assert.Equal(t, 0, list.Len())
    }
}

func TestListImpl_PushFront_And_Remove(t *testing.T) {
    list := New()
    defer cleanup(list)

    list.PushFront(testData1)
    list.PushFront(testData2)

    if assert.NotNil(t, list) {
        assert.Equal(t, 2, list.Len())
        assert.False(t, list.First() == list.Last())
        list.RemoveNode(list.Last())
        assert.Equal(t, 1, list.Len())
    }
}

func TestListImpl_PushFront(t *testing.T) {
    list := New()
    defer cleanup(list)

    list.PushFront(testData1)
    list.PushFront(testData2)
    list.PushBack(testData3)
    list.PushBack(testData4)

    if assert.NotNil(t, list) {
        assert.Equal(t, 4, list.Len())
        assert.False(t, list.First() == list.Last())
        assert.Equal(t, testData3, list.First().Next().Next().Data())
        // Remove the second element and see what happens
        list.RemoveNode(list.First().Next().Next())
        assert.Equal(t, testData4, list.First().Next().Next().Data())
        assert.Equal(t, 3, list.Len())
    }
}

/**
 * TC scenario:
 * 1) Initialise an instance of DLL
 * 1) Add two elements
 * 3) Remove elements one by one and check the list length
 * 4) When all the elements are removed, check the list attributes -> head-nil, tail-nil
 * 5) Repeat step 2) and check the list
 */
func TestListImpl_Repeatability(t *testing.T) {
    // 1)
    list := New()
    defer cleanup(list)
    // 2)
    assert.Equal(t, 0, list.Len())
    list.PushFront(testData1)
    assert.Equal(t, 1, list.Len())
    list.PushBack(testData2)
    assert.Equal(t, 2, list.Len())
    // 3)
    list.RemoveNode(list.First())
    assert.Equal(t, 1, list.Len())
    list.RemoveNode(list.Last())
    assert.Equal(t, 0, list.Len())
    // 4)
    if assert.NotNil(t, list) {
        assert.Nil(t, list.First())
        assert.Nil(t, list.Last())
    }
    // 5)
    list.PushFront(testData3)
    assert.Equal(t, 1, list.Len())
    list.PushBack(testData4)
    assert.Equal(t, 2, list.Len())
    assert.Equal(t, testData3, list.First().Data())
    assert.Equal(t, testData4, list.Last().Data())
}



func Test(t *testing.T) {
    list := New(1, 2, 3, 4)
    list.Add(5, 6, 7, 3)

    list.ForEach()
    println("-----------------------")
    fmt.Printf("%v\n", list.Front())
    fmt.Printf("%v\n", list.Back())

    println("-----------------------")
    list.PushFront(0)
    list.PushBack(8)
    fmt.Printf("%v\n", list.Front())
    fmt.Printf("%v\n", list.Back())
    println("-----------------------")
    list.ForEach()
    println("-----------------------")

    list.Remove(7)
    list.Remove(0)
    list.ForEach()

    println("-----------------------")
    node := list.Find(4)
    fmt.Printf("%v\n", node)
    println("-----------------------")

    list.MoveToFront(node)
    list.ForEach()

    println("-----------------------")
    node = list.Find(5)
    fmt.Printf("%v\n", node)
    println("-----------------------")
    list.MoveToBack(node)
    list.ForEach()

    list.Remove(5)
    println("-----------------------")
    list.ForEach()
    println("-----------------------")

    fmt.Printf("%v\n", list.FindAll(3)) 
    
    println("-----------------------")
    list.MoveAfter(list.Find(4), list.Find(8))
    list.ForEach()
    println("-----------------------")
    fmt.Printf("%v\n", list.Back())
    
    println("-----------------------")
    list.MoveBefore(list.Find(8), list.Front())
    list.ForEach()
    println("-----------------------")
    fmt.Printf("%v\n", list.Front())
      
    println("-----------------------")
    list.InsertBefore(100, list.Front())
    list.ForEach()
    println("-----------------------")
    fmt.Printf("%v\n", list.Front())
      
    println("-----------------------")
    list.InsertAfter(200, list.Front())
    list.InsertAfter(200, list.Front())
    list.ForEach()
    println("-----------------------")
    fmt.Printf("%v\n", list.Front()) 
    
    println("-----------------------")
    list.InsertAfter(300, list.Back())
    list.ForEach()
    println("-----------------------")
    fmt.Printf("%v\n", list.Back()) 
    
    println("----------x-------------")
    list.Reverse()
    list.ForEach() 
    println("-----------------------")
    fmt.Printf("%v\n", list.Front())
    println("-----------------------")
    list = New(1, 2, 3, 4,5,6,7,8,9)
    list.InsertAfter(200, list.Front())
    list.ForEach() 
    println("-----------------------")
    list.Reverse()
    list.ForEach() 
    
}