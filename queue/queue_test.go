package queue

import (
	"testing"
)

func TestNew(t *testing.T) {

	got := New()

	if got == nil || got == new(Queue) {
		t.Errorf("queue.New() = %v; want %v", got, New())
	}
}

func TestDequeue(t *testing.T) {

	t.Run("Pop for empty queue", func(t *testing.T) {

		queue := New()
		got, err := queue.Dequeue()

		if got != nil && err != ErrQueueIsEmpty {
			t.Errorf("queue.Dequeue() = %v, %v; want %v, %v", got, err, nil, ErrQueueIsEmpty)
		}

	})

	t.Run("Pop for non empty queue", func(t *testing.T) {

		queue := New()
		queue.Enqueue(1)
		got, err := queue.Dequeue()

		if err != nil && got != 1 {
			t.Errorf("queue.Dequeue() = %v, %v; want %v, %v", got, err, 1, nil)
		}
	})
}

func TestEnqueue(t *testing.T) {

	queue := New()
	for want := 1; want < 1000; want++ {
		queue.Enqueue(want)
		got := queue.Size()
		if got != want {
			t.Errorf("queue.Enqueue() = %v; want %v", got, want)
		}
	}
}

func TestFront(t *testing.T) {

	queue := New()
	queue.Enqueue(1)
	queue.Enqueue(2)
	got, _ := queue.Front()
	want := 1

	if got != want {
		t.Errorf("queue.Front() = %v; want %v", got, want)
	}
}

func TestEmpty(t *testing.T) {

	t.Run("Pop for empty queue", func(t *testing.T) {
		queue := New()
		got := queue.IsEmpty()

		if !got {
			t.Errorf("queue.IsEmpty() = %v; want %v", got, true)
		}
	})

	t.Run("Pop for non empty queue", func(t *testing.T) {
		queue := New()
		queue.Enqueue(1)
		got := queue.IsEmpty()

		if got {
			t.Errorf("queue.IsEmpty() = %v; want %v", got, false)
		}
	})
}

func TestSize(t *testing.T) {

	queue := New()
	for want := 1; want < 1000; want++ {
		queue.Enqueue(want)
		got := queue.Size()

		if got != want {
			t.Errorf("queue.Size() = %v; want %v", got, want)
		}
	}
}

func TestString(t *testing.T) {

	queue := New()
	queue.Enqueue(1)
	queue.Enqueue(2)
	got := queue.String()
	want := "Queue(1,2)"

	if got != want {
		t.Errorf("queue.String() = %v; want %v", got, want)
	}

}


