package stack

import (
    "fmt"
    "strconv"
    "strings"
)

type Stack struct {
    data []interface{}
    size int
}

func New(initCap ...int) *Stack {
    var capacity int = 10
    if len(initCap) > 0 {
        capacity = initCap[0]
    }

    return &Stack{data: make([]interface{}, 0, capacity)}

}

func (s *Stack) Empty() bool {
    return s.size == 0
}

func (s *Stack) Size() int {
    return s.size
}

func (s *Stack) Push(v interface{}) {
    s.data = append(s.data, v)
    s.size++
}

func (s *Stack) Top() (top interface{}, err error) {
    if s.size > 0 {
        top = s.data[s.size - 1]
        return
    }
    err = fmt.Errorf("Top Error: Stack is empty")
    return
}

func (s *Stack) Pop() (top interface{}, err error) {
    top, err = s.Top();
    if err == nil {
        n := s.size - 1
        s.data[n] = nil
        s.data = s.data[:n]
        s.size--
        return
    }
    err = fmt.Errorf("Top Error: Stack is empty")
    return 
}

func (s *Stack) String() string {
    ret := make([]string, 0, s.size)
    for _, i := range s.data {
        switch i.(type) {
        case string:
            ret = append(ret, i.(string))
        case int:
            ret = append(ret, strconv.Itoa(i.(int)))
        case int32:
            ret = append(ret, strconv.FormatInt(int64(i.(int32)), 10))
        case int64:
            ret = append(ret, strconv.FormatInt(i.(int64), 10))

        case float64:
            ret = append(ret, strconv.FormatFloat(i.(float64), 'f', -1, 64))

        }
    }
    return "[" + strings.Join(ret, ", ") + "]"
}

