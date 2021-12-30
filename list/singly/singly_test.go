package singlylinkedlist

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {

	got := New()

	if got == nil {
		t.Errorf("list.New() = %v; want %v", got, New())
	}
}

func TestEmpty(t *testing.T) {

	list := New()
	got := list.IsEmpty()

	t.Run("For empty list", func(t *testing.T) {
		if got != true {
			t.Errorf("list.IsEmpty() = %v; want %v", got, true)
		}
	})

	t.Run("For non empty list", func(t *testing.T) {

		list = New(1, 2, 3)
		list.Clear()
		got = list.IsEmpty()

		if got != true {
			t.Errorf("list.IsEmpty() = %v; want %v", got, true)
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
  

func TestPopFrontAndFirst(t *testing.T) {

	list := New()
	list.Add(1, 2, 3)

	node := list.PopFront()
	got := node.value
    want := 1

	if got != want {
		t.Errorf("list.PopFront() = %v; want %v", got, want)
	}
    got = list.First().Value()
    want = 2
    if got != want {
		t.Errorf("list.First() = %v; want %v", got, want)
	}

}


func TestFind(t *testing.T) {

	list := New()

	list.Add(0, 1, 2, 3)
	for i := 0; i < list.Len(); i++ {
		node := list.Find(i)
		if node == nil {
			t.Fatalf("list.Find(%d) = Node not found", i)
		}
		got := node.Value()
		want := i

		if got != want {
			t.Errorf("list.Find(%d) = %v; want %v", got, want, i)
		}

	}
}

func TestGet(t *testing.T) {

	list := New()

	list.Add(0, 1, 2, 3)
	for i := 0; i < list.Len(); i++ {
		node := list.Get(i)
		if node == nil {
			t.Fatalf("list.Get(%d) = Node not found", i)
		}
		got := node.Value()
		want := i

		if got != want {
			t.Errorf("list.Get(%d) = %v; want %v", got, want, i)
		}

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

func TestReverse(t *testing.T) {

	got := New()
	got.Add(1, 2, 3)
	got.Reverse()
	want := New()
	want.Add(3, 2, 1)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("list.Reverse() = %v; want %v", got, want)
	}
}
