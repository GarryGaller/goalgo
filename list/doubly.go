package doublylinkedlist

import (
    "fmt"
)

type Node struct {
    data interface{}
    prev *Node
    next *Node
}

type List struct {
    head *Node
    tail *Node
    size int
}

func (n *Node) String() string {
    return fmt.Sprintf(
        "Node(data=%v)", n.data,
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

func (l *List) ForEach() {
    for n := l.head; n != nil; n = n.next {
        fmt.Printf("%v\n", n)
    }
}

func New(values ...interface{}) *List {
    list := &List{}
    if len(values) > 0 {
        list.Add(values...)
    }
    return list
}

func (l *List) Len() int {
    return l.size
}

func (l *List) Empty() bool {
    return l.size == 0
}

func (l *List) Clear() {
    l.size = 0
    l.head = nil
    l.tail = nil
}

func (l *List) Add(values ...interface{}) {

    for _, value := range values {
        l.PushBack(value)
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

func (l *List) PushBack(data interface{}) *Node {

    if data == nil {
        return nil
    }

    node := &Node{data: data}

    l.push(node, "back") 
    return node
}

func (l *List) PushFront(data interface{}) *Node {
    if data == nil {
        return nil
    }

    node := &Node{data: data}

    l.push(node, "front")
    return node
}

func (l *List) InsertAfter(data interface{}, mark *Node) *Node {
    node := &Node{data:data}
    l.insert(node, mark, "after")
    return node
}

func (l *List) InsertBefore(data interface{}, mark *Node) *Node {

    node := &Node{data:data}
    l.insert(node, mark, "before")
    return node
}

func (l *List) MoveAfter(node, mark *Node) {

    l.move(node, mark, "after")
}

func (l *List) MoveBefore(node, mark *Node) {

    l.move(node, mark, "before")
}

func (l *List) MoveToBack(node *Node) {

    l.unlink(node)
    node.prev = l.tail
    node.next = nil
    l.push(node, "back")

}

func (l *List) MoveToFront(node *Node) {

    l.unlink(node)
    node.next = l.head
    node.prev = nil
    l.push(node, "front")

}

func (l *List) Remove(v interface{}) *Node{
    for n := l.head; n != nil; n = n.next {
        if n.data == v {
            l.remove(n)
            return n
        }
    } 
    return nil
}

func (l *List) RemoveAll(v interface{}) int {
    var count int
    for n := l.head; n != nil; n = n.next {
        if n.data == v {
            l.remove(n)
            count++
        }
    }
    return count
}
   

func (l *List) RemoveNode(node *Node) {
    l.remove(node)
}


func (l *List) Reverse2() {
   var prev *Node
   var current *Node = l.head
   l.tail = l.head
   
   for current != nil {
        prev = current.prev
        current.prev = current.next
        current.next = prev          
        current = current.prev
   } 
   
   if prev != nil {
        l.head = prev.prev
   }
}
   
func (l *List) Reverse() {
   var prev *Node
   l.tail = l.head
   
   for  n := l.head; n != nil; n = n.prev{
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

func (l *List) insert(node, mark *Node, position string) {
    
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
}
 

func (l *List) remove(node *Node) {
    l.unlink(node)
    l.delete(node)

}
func (l *List) unlink(node *Node) {

    if node == l.head {
        l.head = node.next
    }
    if node == l.tail {
        l.tail = node.prev
    }

    prev := node.prev
    next := node.next

    if prev != nil {
        prev.next = next
    }
    if next != nil {
        next.prev = prev
    }
}

func (l *List) delete(node *Node) {
    node.next = nil
    node.prev = nil
    l.size--

    if l.size == 0 {
        l.Clear()
    }
}

func (l *List) move(node, mark *Node, position string) {

    if l.size == 0 {
        return  
    }
    
    l.unlink(node)
    l.insert(node, mark, position)
}




 

//         1                            2                  3
//<- prev <=> next->           <- prev <=> next-> <- prev <=> next->


