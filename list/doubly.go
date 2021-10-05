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

func (l *List) PushBack(data interface{}) (out *Node) {

	if data == nil {
		return nil
	}

	out = &Node{data: data}

	l.push(out, "back")
	return
}

func (l *List) PushFront(data interface{}) (out *Node) {

	if data == nil {
		return nil
	}

	out = &Node{data: data}

	l.push(out, "front")
	return
}

func (l *List) InsertAfter(data interface{}, mark *Node) (out *Node) {

	if mark == nil {
		return nil
	}

	node := &Node{data: data}
	if l.insert(node, mark, "after") {
		out = node
	}

	return
}

func (l *List) InsertBefore(data interface{}, mark *Node) (out *Node) {

	if mark == nil {
		return nil
	}

	node := &Node{data: data}
	if l.insert(node, mark, "before") {
		out = node
	}

	return
}

func (l *List) MoveAfter(node, mark *Node) bool {

	return l.move(node, mark, "after")
}

func (l *List) MoveBefore(node, mark *Node) bool {

	return l.move(node, mark, "before")
}

func (l *List) MoveToBack(node *Node) bool {

	if node == nil {
		return false
	}

	l.unlink(node)
	node.prev = l.tail
	node.next = nil
	l.push(node, "back")

	return true
}

func (l *List) MoveToFront(node *Node) bool {

	if node == nil {
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

// not work !!!
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
	if node != nil {
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

	if node == nil {
		return
	}

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

	if mark == nil {
		return false
	}

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
	if node != nil {
		l.unlink(node)
		l.delete(node)
	}
}

func (l *List) unlink(node *Node) {

	if node == nil {
		return
	}
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

	if node != nil {
		node.next = nil
		node.prev = nil
		l.size--

		if l.size == 0 {
			l.Clear()
		}

		node = nil
	}
}

func (l *List) move(node, mark *Node, position string) (out bool) {

	if l.size != 0 && node != nil && mark != nil {
		l.unlink(node)
		out = l.insert(node, mark, position)
	}

	return
}

//         1                            2                  3
//<- prev <=> next->           <- prev <=> next-> <- prev <=> next->
