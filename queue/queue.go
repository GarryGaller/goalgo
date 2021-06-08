package queue

import (
    "strconv"
    "strings"
    "errors"
)

var errorEmpty = errors.New("Front Error: Queue is empty")


type  Queue struct {
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

func (q *Queue) Empty() bool {
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
    err = errorEmpty
    return 
}


func (q *Queue) Dequeue() (top interface{}, err error) {
    top, err = q.Front()
    if  err == nil {
        q.data[0] = nil
        q.data = q.data[1:]
        q.size--
        return 
    }
    err = errorEmpty
    return 
}
 

func (q *Queue) String() string {
    ret := make([]string, 0, q.size)
    for _, i := range q.data {
        switch i.(type) {
        case string:
            ret = append(ret, i.(string))
        case int:
            ret = append(ret, strconv.Itoa(i.(int)))
        case int32:
            ret = append(ret, strconv.FormatInt(int64(i.(int32)), 10))
        case int64:
            ret = append(ret, strconv.FormatInt(i.(int64),10))    
        
        case float64:
            ret = append(ret, strconv.FormatFloat(i.(float64), 'f', -1, 64))
        
        }
    }
    return "[" + strings.Join(ret,", ") + "]"
}



