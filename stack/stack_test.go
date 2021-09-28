package stack

import (
	//"reflect"
	"testing"
)

func TestNew(t *testing.T) {

	got := New()

	if got == nil || got == new(Stack) {
		t.Errorf("stack.New() = %v; want %v", got, New())
	}
}

func TestPop(t *testing.T) {

	t.Run("Pop for empty stack", func(t *testing.T) {

		stack := New()
		got, err := stack.Pop()

		if got != nil && err != ErrStackIsEmpty {
			t.Errorf("stack.Pop() = %v, %v; want %v, %v", got, err, nil, ErrStackIsEmpty)
		}

	})

	t.Run("Pop for non empty stack", func(t *testing.T) {

		stack := New()
		stack.Push(1)
		got, err := stack.Pop()

		if err != nil && got != 1 {
			t.Errorf("stack.Pop() = %v, %v; want %v, %v", got, err, 1, nil)
		}
	})
}

func TestPush(t *testing.T) {

	stack := New()
	for want := 1; want < 1000; want++ {
		stack.Push(want)
		got := stack.Size()
		if got != want {
			t.Errorf("stack.Push() = %v; want %v", got, want)
		}
	}
}

func TestTop(t *testing.T) {

	stack := New()
	for want := 1; want < 1000; want++ {
		stack.Push(want)
		got, _ := stack.Top()
		if got != want {
			t.Errorf("stack.Pop() = %v; want %v", got, want)
		}
	}
}

func TestEmpty(t *testing.T) {

	t.Run("Pop for empty stack", func(t *testing.T) {
		stack := New()
		got := stack.IsEmpty()

		if !got {
			t.Errorf("stack.IsEmpty() = %v; want %v", got, true)
		}
	})

	t.Run("Pop for non empty stack", func(t *testing.T) {
		stack := New()
		stack.Push(1)
		got := stack.IsEmpty()

		if got {
			t.Errorf("stack.IsEmpty() = %v; want %v", got, false)
		}
	})
}

func TestSize(t *testing.T) {

	stack := New()
	for want := 1; want < 1000; want++ {
		stack.Push(want)
		got := stack.Size()

		if got != want {
			t.Errorf("stack.Size() = %v; want %v", got, want)
		}
	}
}

func TestString(t *testing.T) {

	stack := New()
	stack.Push(1)
	stack.Push(2)
	got := stack.String()
	want := "Stack(1,2)"

	if got != want {
		t.Errorf("stack.String() = %v; want %v", got, want)
	}

}
