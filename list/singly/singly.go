package singlylinkedlist

type Node struct {
    value interface{}
    next  *Node
}

type LinkedList struct {
    head *Node
    size int
}

func New(values ...interface{}) *LinkedList {
    list := &LinkedList{}
    if len(values) > 0 {
        list.Add(values...)
    }
    return list
}

func (n *Node) Next() *Node {
    return n.next
}

func (n *Node) Value() interface{} {
    return n.value
}

func (l *LinkedList) Add(values ...interface{}) {

    for _, value := range values {
        l.PushBack(value)
    }
}

func (l *LinkedList) Append(values ...interface{}) {
    l.Add(values...)
}

func (l *LinkedList) First() *Node {
    return l.head
}

func (l *LinkedList) Front() *Node {
    return l.head
}

func (l *LinkedList) Get(index int) *Node {
    idx := 0
    for n := l.head; n != nil; n = n.next {
        if idx == index {
            return n
        }
        idx++
    }

    return nil
}

func (l *LinkedList) Find(v interface{}) *Node {
    for n := l.head; n != nil; n = n.next {
        if n.value == v {
            return n
        }
    }
    return nil
}

func (l *LinkedList) Len() int {
    return l.size
}

func (l *LinkedList) IsEmpty() bool {
    return l.size == 0
}

func (l *LinkedList) Clear() {

    l.size = 0
    l.head = nil
}

func (l *LinkedList) Has(v interface{}) bool {

    return l.Find(v) != nil
}

func (l *LinkedList) PushBack(value interface{}) {
    node := &Node{value:value}
   
    if l.size == 0 {
        l.head = node
        l.size++
        return
    }
    head := l.head
    for i := 0; i < l.size; i++ {
        if head.next == nil {
            head.next = node
            l.size++
            return
        }
        head = head.next
    }
}
 
func (l *LinkedList) PushFront(value interface{}) {
     node := &Node{value:value}
     node.next = l.head
     l.head = node
     l.size++
}


func (l *LinkedList) PopFront() *Node {
    head := l.head
    l.head = head.next
    return head
}

func (l *LinkedList) Reverse() {

    var prev, next *Node
    var curr = l.head

    for curr != nil {
        next = curr.next
        curr.next = prev
        prev = curr
        curr = next
    }
    l.head = prev
}

//TODO:
//IndexOf(value)
//Remove(index)
//Remove(node)
//Swap
//Set(index, value)
