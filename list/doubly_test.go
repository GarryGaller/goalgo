package doublylinkedlist

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleForeach() {
	list := New(1, 2, 3)
	list.ForEach()
	// Output:
	// Node(data=1)
	// Node(data=2)
	// Node(data=3)

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
	// Node(data=2)
	// Node(data=4)
}

func ExampleRemove2() {

	list := New()
	list.Add(3, 3, 3)
	removed := list.Remove(3)
	fmt.Println(removed)
    list.ForEach()
	// Output:
    // true
	// Node(data=3)
	// Node(data=3)
}
  

func ExampleRemoveNode() {
	list := New(1, 2, 3, 4)
	list.RemoveNode(list.First())
	list.ForEach()
	// Output:
	// Node(data=2)
    // Node(data=3)
	// Node(data=4)
}

  
func ExampleRemoveAll() {

	list := New()
	list.Add(3, 3, 3)
	list.RemoveAll(3)
	list.ForEach()
	// Output:
	//
}
  
func ExampleRemoveAll2() {

	list := New()
	list.Add(1, 3, 3)
	list.RemoveAll(3)
	list.ForEach()
	// Output:
	// Node(data=1)
}


func ExampleInsertBefore() {
	list := New(1, 2, 3)
	list.InsertBefore(0, list.First())
	list.ForEach()
	// Output:
	// Node(data=0)
	// Node(data=1)
	// Node(data=2)
	// Node(data=3)

}

func ExampleInsertAfter() {
	list := New(0, 2, 3)
	list.InsertAfter(1, list.First())
	list.ForEach()
	// Output:
	// Node(data=0)
	// Node(data=1)
	// Node(data=2)
	// Node(data=3)

}

func ExampleMoveToFront() {
	list := New(1, 2, 3)
	node := list.Find(3)
	moved := list.MoveToFront(node)
	fmt.Println(moved)
	list.ForEach()
	// Output:
	// true
	// Node(data=3)
	// Node(data=1)
	// Node(data=2)
}

func ExampleMoveToBack() {
	list := New(1, 2, 3)
	node := list.Find(1)
	moved := list.MoveToBack(node)
	fmt.Println(moved)
	list.ForEach()
	// Output:
	// true
	// Node(data=2)
	// Node(data=3)
	// Node(data=1)
}

func ExampleMoveAfter() {
	list := New(1, 2, 3)
	node := list.Find(1)
	moved := list.MoveAfter(node, list.Back())
	fmt.Println(moved)
	list.ForEach()
	// Output:
	// true
	// Node(data=2)
	// Node(data=3)
	// Node(data=1)
}

func ExampleMoveBefore() {
	list := New(1, 2, 3)
	node := list.Find(1)
	moved := list.MoveBefore(node, list.Back())
	fmt.Println(moved)
	list.ForEach()
	// Output:
	// true
	// Node(data=2)
	// Node(data=1)
	// Node(data=3)
}

func ExampleReverse() {

	list := New()
	list.Add(1, 2, 3)
	list.Reverse()
	list.ForEach()
	// Output:
	// Node(data=3)
	// Node(data=2)
	// Node(data=1)
}




func TestNew(t *testing.T) {

	got := New()

	if got == nil || got == new(List) {
		t.Errorf("list.New() = %v; want %v", got, New())
	}
}

func TestEmpty(t *testing.T) {

	list := New()
	got := list.Empty()

	t.Run("For empty list", func(t *testing.T) {
		if got != true {
			t.Errorf("list.Empty() = %v; want %v", got, true)
		}
	})

	t.Run("For non empty list", func(t *testing.T) {

		list = New(1, 2, 3)
		list.Clear()
		got = list.Empty()

		if got != true {
			t.Errorf("list.Empty() = %v; want %v", got, true)
		}

	})

}

func TestAdd(t *testing.T) {

	list := New()
	list.Add(1, 2, 3)

	got := list.Len()
	want := 3

	if got != want {
		t.Errorf("list.Len() = %v; want %v", got, want)
	}

}

func TestFrontAndPushFront(t *testing.T) {

	list := New()

	for want := 10; want >= 0; want-- {
		list.PushFront(want)
		node := list.Front()
		got := node.Value()
		if got != want {
			t.Errorf("list.Front() = %v; want %v", got, want)
		}
	}
}

func TestBackAndPushBack(t *testing.T) {

	list := New()

	for want := 0; want <= 10; want++ {
		list.PushBack(want)
		node := list.Back()
		got := node.Value()
		if got != want {
			t.Errorf("list.Back() = %v; want %v", got, want)
		}
	}
}

func TestInsertBefore(t *testing.T) {

	list := New(1, 2, 3)
	got := list.InsertBefore(0, list.Front())
	want := list.Front()

	if got == nil || !reflect.DeepEqual(got, want) {
		t.Errorf("list.InsertBefore() = %v; want %v", got, want)
	}

	gotValue := got.Value()
	if gotValue != 0 {
		t.Errorf("list.InsertBefore().Value(): %v; want %v", gotValue, 1)
	}

	wantLen := 4

	if list.Len() != wantLen {
		t.Errorf("list.Len() = %v; want %v", list.Len(), wantLen)
	}

}

func TestInsertAfter(t *testing.T) {

	list := New(0, 2, 3)
	got := list.InsertAfter(1, list.Front())
	want := list.Front().Next()

	if got == nil || !reflect.DeepEqual(got, want) {
		t.Errorf("list.InsertAfter() = %v; want %v", got, want)
	}

	gotValue := got.Value()
	if gotValue != 1 {
		t.Errorf("list.InsertAfter().Value(): %v; want %v", gotValue, 1)
	}

	wantLen := 4

	if list.Len() != wantLen {
		t.Errorf("list.Len() = %v; want %v", list.Len(), wantLen)
	}

}

func TestLen(t *testing.T) {

	list := New()
	list.Add(1, 2, 3)

	got := list.Len()
	want := 3
	if got != 3 {
		t.Errorf("list.Len() = %v; want %v", got, want)
	}
}

func TestFind(t *testing.T) {

	list := New()
	list.Add(1, 2, 3)

	got := list.Find(3)
	want := list.Back()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("llist.Find = %v; want %v", got, want)
	}
}

func TestFindAll(t *testing.T) {

	list := New()
	list.Add(3, 3, 3)

	got := len(list.FindAll(3))
	want := 3

	if got != want {
		t.Errorf("len(list.FindAll()) = %v; want %v", got, want)
	}
}

func TestRemoveNode(t *testing.T) {

	list := New()
	list.Add(1, 2, 3)
	want := list.Front().Next()
	list.RemoveNode(list.Front())
	got := list.Front()

	if got != want {
		t.Errorf("list.RemoveNode() = %v; want %v", got, want)
	}

	if got.Value() != 2 {
		t.Errorf("list.Front().Value() = %v; want %v", got.Value(), 2)
	}

	gotLen := list.Len()
	wantlen := 2
	if gotLen != wantlen {
		t.Errorf("list.Len() = %v; want %v", gotLen, wantlen)
	}
}


func TestRemoveAll(t *testing.T) {

	list := New()
	list.Add(3, 3, 3)

	wantCount := len(list.FindAll(3))
	got := list.RemoveAll(3)

	if got != wantCount {
		t.Errorf("list.RemoveAll() = %v; want %v", got, wantCount)
	}

	gotLen := list.Len()
	if gotLen != 0 {
		t.Errorf("list.Len() = %v; want %v", gotLen, 0)
	}
}

func TestClear(t *testing.T) {

	list := New()
	list.Add(1, 2, 3)

	list.Clear()
	if !reflect.DeepEqual(list, New()) {
		t.Errorf("list.Clear() = %v; want %v", list, New())
	}

}
