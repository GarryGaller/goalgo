package singlylinkedlist

import (
    "reflect"
    "testing"
)

func TestNew(t *testing.T) {

    got := New()

    if got == nil || got == new(LinkedList) {
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

func TestFind(t *testing.T) {

    list := New()
    list.Add(1, 2, 3)

    got := list.Find(3)
    want := &Node{value:3}

    if !reflect.DeepEqual(got, want) {
        t.Errorf("list.Find() = %v; want %v", got, want)
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
    
    if !reflect.DeepEqual(got, want)  {
        t.Errorf("list.Reverse() = %v; want %v", got, want)
    }
}






