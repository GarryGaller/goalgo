package queue

import (
	"errors"
	"fmt"
	"strings"
)

var ErrQueueIsEmpty = errors.New("Queue is empty")

type Queue struct {
	data []interface{}
	size int
}

func New(initCap ...int) *Queue {
	var capacity int = 10
	if len(initCap) > 0 {
		capacity = initCap[0]
	}

	return &Queue{data: make([]interface{}, 0, capacity)}

}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Enqueue(v interface{}) {
	q.data = append(q.data, v)
	q.size++
}

func (q *Queue) Front() (top interface{}, err error) {
	if q.size > 0 {
		top = q.data[0]
		return
	}
	err = ErrQueueIsEmpty
	return
}

func (q *Queue) Dequeue() (top interface{}, err error) {
	top, err = q.Front()
	if err == nil {
		q.data[0] = nil
		q.data = q.data[1:]
		q.size--
		return
	}
	err = ErrQueueIsEmpty
	return
}

func (q *Queue) String() (out string) {

	var builder strings.Builder
	builder.WriteString("[")

	for _, value := range q.data {
		builder.WriteString(fmt.Sprintf("%#v", value))
		builder.WriteString(", ")
	}
	out = strings.TrimSuffix(builder.String(), ", ") + "]"

	return
}
