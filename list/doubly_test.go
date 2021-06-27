package doublylinkedlist_test

//https://github.com/omerkaya1/doublylinkedlist/blob/master/doublylinkedlist.go

import (
	"fmt"
	"github.com/GarryGaller/goalgo/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testData1 = "first"
	testData2 = "second"
	testData3 = "third"
	testData4 = 1
	testData5 = 2
	testData6 = 3
)

func cleanup(impl *doublylinkedlist.List) {
	impl = nil
	return
}

/**
 * A set o general tests
 */
func TestNewList(t *testing.T) {
	dl := doublylinkedlist.New()
	defer cleanup(dl)

	if assert.NotNil(t, dl) {
		assert.Equal(t, 0, dl.Len())
		assert.Nil(t, dl.First())
		assert.Nil(t, dl.Last())
	}
}

func TestListGeneral(t *testing.T) {
	dl := doublylinkedlist.New()
	defer cleanup(dl)

	dl.PushFront(testData1)
	dl.PushFront(testData2)
	dl.PushFront(testData3)

	dl.PushBack(testData4)
	dl.PushBack(testData5)
	dl.PushBack(testData6)

	if assert.NotNil(t, dl) {
		assert.Equal(t, 6, dl.Len())
		assert.Equal(t, testData3, dl.First().Data())
		assert.Equal(t, testData6, dl.Last().Data())
	}
}

/**
 * A set of test for the doublylinkedlist functionality
 */
func TestListImpl_SingleElementList(t *testing.T) {
	dl := doublylinkedlist.New()
	defer cleanup(dl)

	dl.PushFront(testData1)

	if assert.NotNil(t, dl) {
		assert.Equal(t, 1, dl.Len())
		assert.True(t, dl.First() == dl.Last())
	}
}

func TestListImpl_RemoveSingleElement(t *testing.T) {
	dl := doublylinkedlist.New()
	defer cleanup(dl)

	dl.PushFront(testData1)

	if assert.NotNil(t, dl) {
		assert.Equal(t, 1, dl.Len())
		assert.True(t, dl.First() == dl.Last())
		dl.RemoveNode(dl.First())
		assert.Equal(t, 0, dl.Len())
	}
}

func TestListImpl_PushFront_And_Remove(t *testing.T) {
	dl := doublylinkedlist.New()
	defer cleanup(dl)

	dl.PushFront(testData1)
	dl.PushFront(testData2)

	if assert.NotNil(t, dl) {
		assert.Equal(t, 2, dl.Len())
		assert.False(t, dl.First() == dl.Last())
		dl.RemoveNode(dl.Last())
		assert.Equal(t, 1, dl.Len())
	}
}

func TestListImpl_PushFront(t *testing.T) {
	dl := doublylinkedlist.New()
	defer cleanup(dl)

	dl.PushFront(testData1)
	dl.PushFront(testData2)
	dl.PushBack(testData3)
	dl.PushBack(testData4)

	if assert.NotNil(t, dl) {
		assert.Equal(t, 4, dl.Len())
		assert.False(t, dl.First() == dl.Last())
		assert.Equal(t, testData3, dl.First().Next().Next().Data())
		// Remove the second element and see what happens
		dl.RemoveNode(dl.First().Next().Next())
		assert.Equal(t, testData4, dl.First().Next().Next().Data())
		assert.Equal(t, 3, dl.Len())
	}
}

/**
 * TC scenario:
 * 1) Initialise an instance of DLL
 * 1) Add two elements
 * 3) Remove elements one by one and check the doublylinkedlist length
 * 4) When all the elements are removed, check the doublylinkedlist attributes -> head-nil, tail-nil
 * 5) Repeat step 2) and check the doublylinkedlist
 */
func TestListImpl_Repeatability(t *testing.T) {
	// 1)
	dl := doublylinkedlist.New()
	defer cleanup(dl)
	// 2)
	assert.Equal(t, 0, dl.Len())
	dl.PushFront(testData1)
	assert.Equal(t, 1, dl.Len())
	dl.PushBack(testData2)
	assert.Equal(t, 2, dl.Len())
	// 3)
	dl.RemoveNode(dl.First())
	assert.Equal(t, 1, dl.Len())
	dl.RemoveNode(dl.Last())
	assert.Equal(t, 0, dl.Len())
	// 4)
	if assert.NotNil(t, dl) {
		assert.Nil(t, dl.First())
		assert.Nil(t, dl.Last())
	}
	// 5)
	dl.PushFront(testData3)
	assert.Equal(t, 1, dl.Len())
	dl.PushBack(testData4)
	assert.Equal(t, 2, dl.Len())
	assert.Equal(t, testData3, dl.First().Data())
	assert.Equal(t, testData4, dl.Last().Data())
}

func TestList(t *testing.T) {
	dl := doublylinkedlist.New(1, 2, 3, 4)
	dl.Add(5, 6, 7, 3)

	dl.ForEach()
	println("-----------------------")
	fmt.Printf("%v\n", dl.Front())
	fmt.Printf("%v\n", dl.Back())

	println("-----------------------")
	dl.PushFront(0)
	dl.PushBack(8)
	fmt.Printf("%v\n", dl.Front())
	fmt.Printf("%v\n", dl.Back())
	println("-----------------------")
	dl.ForEach()
	println("-----------------------")

	dl.Remove(7)
	dl.Remove(0)
	dl.ForEach()

	println("-----------------------")
	node := dl.Find(4)
	fmt.Printf("%v\n", node)
	println("-----------------------")

	dl.MoveToFront(node)
	dl.ForEach()

	println("-----------------------")
	node = dl.Find(5)
	fmt.Printf("%v\n", node)
	println("-----------------------")
	dl.MoveToBack(node)
	dl.ForEach()

	dl.Remove(5)
	println("-----------------------")
	dl.ForEach()
	println("-----------------------")

	fmt.Printf("%v\n", dl.FindAll(3))

	println("-----------------------")
	dl.MoveAfter(dl.Find(4), dl.Find(8))
	dl.ForEach()
	println("-----------------------")
	fmt.Printf("%v\n", dl.Back())

	println("-----------------------")
	dl.MoveBefore(dl.Find(8), dl.Front())
	dl.ForEach()
	println("-----------------------")
	fmt.Printf("%v\n", dl.Front())

	println("-----------------------")
	dl.InsertBefore(100, dl.Front())
	dl.ForEach()
	println("-----------------------")
	fmt.Printf("%v\n", dl.Front())

	println("-----------------------")
	dl.InsertAfter(200, dl.Front())
	dl.InsertAfter(200, dl.Front())
	dl.ForEach()
	println("-----------------------")
	fmt.Printf("%v\n", dl.Front())

	println("-----------------------")
	dl.InsertAfter(300, dl.Back())
	dl.ForEach()
	println("-----------------------")
	fmt.Printf("%v\n", dl.Back())

	println("----------x-------------")
	dl.Reverse()
	dl.ForEach()
	println("-----------------------")
	fmt.Printf("%v\n", dl.Front())
	println("-----------------------")
	dl = doublylinkedlist.New(1, 2, 3, 4, 5, 6, 7, 8, 9)
	dl.InsertAfter(200, dl.Front())
	dl.ForEach()
	println("-----------------------")
	dl.Reverse()
	dl.ForEach()

}
