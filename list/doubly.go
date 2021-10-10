package doublylinkedlist

import (
    "fmt"
    "strings"
)

type Node struct {
    data interface{}
    prev *Node
    next *Node
    list *List
}

type List struct {
    head *Node
    tail *Node
    size int
}

func (n *Node) String() string {
    return fmt.Sprintf(
        "Node(data=%#v)", n.data,
    )
}

func (n *Node) Value() interface{} {
    return n.data
}

func (n *Node) Data() interface{} {
    return n.data
}

func (n *Node) Next() *Node {
    return n.next
}

func (n *Node) Prev() *Node {
    return n.prev
}

func New(values ...interface{}) *List {
    list := &List{}
    if len(values) > 0 {
        list.Add(values...)
    }
    return list
}

func (l *List) Slice() (out []interface{}) {

    for n := l.Front(); n != nil; n = n.Next() {
        out = append(out, n.Value())
    }

    return
}

func (l *List) Values() (out []interface{}) {

    for n := l.Front(); n != nil; n = n.Next() {
        out = append(out, n.Value())
    }

    return
}

func (l *List) ForEach(f ...func(...interface{}) (int, error)) {

    if len(f) != 0 {
        for n := l.head; n != nil; n = n.next {
            f[0](n)
        }
    } else {
        for n := l.head; n != nil; n = n.next {
            fmt.Printf("%#v\n", n.Value())
        }
    }
}

func (l *List) ForEach2(format ...string) {
    f := ""
    if len(format) > 0 {
        f = format[0]
    }

    switch f {

    case "":
        for n := l.head; n != nil; n = n.next {
            fmt.Printf("%#v\n", n.Value())
        }
    case "%v":
        for n := l.head; n != nil; n = n.next {
            fmt.Printf("%v\n", n)
        }

    case "%+v":
        for n := l.head; n != nil; n = n.next {
            fmt.Printf("%+v\n", n)
        }

    case "%#v":
        for n := l.head; n != nil; n = n.next {
            fmt.Printf("%#v\n", n)
        }
    default:
        fmt.Printf("Invalid format:%s\n", f)
    }
}

func (l *List) String() (out string) {

    var builder strings.Builder
    builder.WriteString("[")

    for n := l.Front(); n != nil; n = n.Next() {
        builder.WriteString(fmt.Sprintf("%#v", n.Value()))
        builder.WriteString(", ")
    }
    out = strings.TrimSuffix(builder.String(), ", ") + "]"

    return
}

func (l *List) Repr() (out string) {
    // TO DO

    return
}

func (l *List) Len() int {
    return l.size
}

func (l *List) IsEmpty() bool {
    return l.size == 0
}

func (l *List) Clear() {

    for n := l.Back(); n != nil; n = n.Prev() {
        n.next = nil
    }

    l.size = 0
    l.head = nil
    l.tail = nil
}
  

func (l *List) Clear2() {

    l.size = 0
    l.head = nil
    l.tail = nil
}


func (l *List) Add(values ...interface{}) {

    for _, value := range values {
        l.PushBack(value)
    }
}

func (l *List) Append(values ...interface{}) {
    l.Add(values...)
}

func (l *List) Prepend(values ...interface{}) {

    at := l.PushFront(values[0])

    for _, value := range values[1:] {
        at = l.InsertAfter(value, at)
    }

}

func (l *List) SwapValue(i, j *Node) bool {

    if i == nil || i.list != l || j == nil || j.list != l {
        return false
    }

    jdata := j.data
    j.data = i.data
    i.data = jdata

    return true

}

func (l *List) ReplaceValue(n *Node, v interface{}) {
    n.data = v
}

func (l *List) PushBackList(other *List) {

    for n := other.Front(); n != nil; n = n.Next() {
        l.PushBack(n.Value())
    }
}

func (l *List) PushFrontList(other *List) {

    for n := other.Back(); n != nil; n = n.Prev() {
        l.PushFront(n.Value())
    }
}

func (l *List) Back() *Node {
    return l.tail
}

func (l *List) Front() *Node {
    return l.head
}

func (l *List) First() *Node {
    return l.head
}

func (l *List) Last() *Node {
    return l.tail
}

func (l *List) Find(v interface{}) *Node {
    for n := l.head; n != nil; n = n.next {
        if n.data == v {
            return n
        }
    }
    return nil
}

func (l *List) FindAll(v interface{}) []*Node {
    nodes := make([]*Node, 0, 1)
    for n := l.head; n != nil; n = n.next {
        if n.data == v {
            nodes = append(nodes, n)
        }
    }
    return nodes
}

func (l *List) Has(v interface{}) bool {

    return l.Find(v) != nil
}

func (l *List) PushBack(data interface{}) (out *Node) {

    if data == nil {
        return nil
    }

    out = &Node{data: data, list: l}

    l.push(out, "back")
    return
}

func (l *List) PushFront(data interface{}) (out *Node) {

    if data == nil {
        return nil
    }

    out = &Node{data: data, list: l}

    l.push(out, "front")
    return
}

func (l *List) InsertAfter(data interface{}, mark *Node) (out *Node) {

    if mark == nil || mark.list != l {
        return nil
    }

    node := &Node{data: data, list: l}
    if l.insert(node, mark, "after") {
        out = node
    }

    return
}

func (l *List) InsertBefore(data interface{}, mark *Node) (out *Node) {

    if mark == nil || mark.list != l {
        return nil
    }

    node := &Node{data: data, list: l}
    if l.insert(node, mark, "before") {
        out = node
    }

    return
}

func (l *List) MoveAfter(node, mark *Node) bool {

    if node == nil || mark == nil || l.size <= 1 {
        return false
    }

    if node.list != l || node == mark || mark.list != l {
        return false
    }

    return l.move(node, mark, "after")
}

func (l *List) MoveBefore(node, mark *Node) bool {

    if node == nil || mark == nil || l.size <= 1 {
        return false
    }

    if node.list != l || node == mark || mark.list != l {
        return false
    }

    return l.move(node, mark, "before")
}

func (l *List) MoveToBack(node *Node) bool {

    if node == nil || node.list != l || node == l.tail {
        return false
    }

    l.unlink(node)
    node.prev = l.tail
    node.next = nil
    l.push(node, "back")

    return true
}

func (l *List) MoveToFront(node *Node) bool {

    if node == nil || node.list != l || node == l.head {
        return false
    }

    l.unlink(node)
    node.next = l.head
    node.prev = nil
    l.push(node, "front")

    return true
}

func (l *List) Remove(v interface{}) bool {
    out := l.Find(v)
    if out != nil {
        l.remove(out)
        return true
    }
    return false
}

func (l *List) RemoveFront() {
    l.remove(l.Front())
}

func (l *List) RemoveBack() {
    l.remove(l.Back())
}

func (l *List) RemoveFirst() {
    l.remove(l.First())
}

func (l *List) RemoveLast() {
    l.remove(l.Last())
}

/*
// incorrect version
func (l *List) Bad_RemoveAll(v interface{}) int {
    var count int
    for n := l.head; n != nil; n = n.next {
        if n.data == v {
            l.remove(n)
            count++
        }
    }
    return count
}
*/

func (l *List) RemoveAll(v interface{}) int {
    var count int
    nodes := l.FindAll(v)
    for _, n := range nodes {
        l.remove(n)
        count++
    }
    return count
}

func (l *List) RemoveNode(node *Node) {
    if node != nil && node.list == l {
        l.remove(node)
    }
}

func (l *List) Reverse2() {
    var prev *Node
    var n *Node = l.head
    l.tail = l.head

    for n != nil {
        prev = n.prev
        n.prev = n.next
        n.next = prev
        n = n.prev
    }

    if prev != nil {
        l.head = prev.prev
    }
}

func (l *List) Reverse() {
    var prev *Node
    l.tail = l.head

    for n := l.head; n != nil; n = n.prev {
        prev = n.prev
        n.prev = n.next
        n.next = prev
    }

    if prev != nil {
        l.head = prev.prev
    }
}

func (l *List) push(node *Node, position string) {

    if l.size == 0 {
        l.head = node
        l.tail = node
    } else {

        switch position {
        case "front":
            l.head.prev = node
            node.next = l.head
            l.head = node

        case "back":
            l.tail.next = node
            node.prev = l.tail
            l.tail = node
        }
    }
    l.size++
}

func (l *List) insert(node, mark *Node, position string) bool {

    switch position {

    case "before":
        if mark == l.head {
            l.head = node
        }

        node.prev = mark.prev
        node.next = mark
        if mark.prev != nil {
            mark.prev.next = node
        }

        mark.prev = node

    case "after":
        if mark == l.tail {
            l.tail = node
        }

        node.prev = mark
        node.next = mark.next
        if mark.next != nil {
            mark.next.prev = node
        }
        mark.next = node
        //
    }
    l.size++
    return true
}

func (l *List) remove(node *Node) {
    l.unlink(node)
    l.delete(node)
}

func (l *List) unlink(node *Node) {

    // если узел - голова
    if node == l.head {
        // говорим что новая голова - следующий за head узел
        l.head = node.next
    }
    // если узел - хвост
    if node == l.tail {
        // говорим что новый хвост - предыдущий перед tail узел
        l.tail = node.prev
    }

    prev := node.prev
    next := node.next
    // если перед удаляемым узлом есть другой
    if prev != nil {
        // переставляем указатель его next на следующий за удаляемым узел
        prev.next = next
    }
    // если после удаляемого узла есть другой
    if next != nil {
        // переставляем указатель его prev на предшественника удаляемого узла
        next.prev = prev
    }
}

func (l *List) delete(node *Node) {

    node.next = nil
    node.prev = nil
    l.size--
    node = nil
}

func (l *List) move(node, mark *Node, position string) (out bool) {

    l.unlink(node)
    out = l.insert(node, mark, position)

    return
}

// list.New
// list.First
// list.Last
// list.Front
// list.Back
// list.Clear
// list.IsEmpty
// list.Len
// list.Add
// list.Append
// list.Prepend
// list.PushBack
// list.PushBackList
// list.PushFront
// list.PushFrontList
// list.RemoveFront
// list.RemoveBack
// list.RemoveFirst
// list.RemoveLast
// list.Remove
// list.RemoveAll
// list.RemoveNode
// list.Reverse
// list.SwapValue
// list.ReplaceValue
// list.Find
// list.FindAll
// list.InsertAfter
// list.InsertBefore
// list.MobeAfter
// list.MoveBefore
// list.MoveToBack
// list.MoveToFirst
// list.Slice
// list.Values
// list.String
// list.ForEach

//              1                                2                        3
// nil <- prev <=> next-> 2           1 <- prev <=> next-> 3   2 <- prev <=> next-> nil
