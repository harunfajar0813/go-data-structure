package linked

import "fmt"

type Node struct {
	value    int
	next     *Node
	previous *Node
}

func NewNode(value int) (error, *Node) {
	return nil, &Node{
		value: value,
	}
}

type DoublyCircularLinked struct {
	head   *Node
	tail   *Node
	length int
}

func (l *DoublyCircularLinked) CircularConnection() {
	l.head.previous = l.tail
	l.head.previous.next = l.head
}

func (l *DoublyCircularLinked) Insert(value int) {
	if err, newNode := NewNode(value); err == nil {
		if l.tail == nil {
			l.tail = newNode
			l.head = l.tail
		} else {
			newNode.next = l.head
			newNode.next.previous = newNode
			l.head = l.head.previous
		}
		l.CircularConnection()
		l.length++
	}
}

func (l *DoublyCircularLinked) Append(value int) {
	if err, newNode := NewNode(value); err == nil {
		if l.tail == nil {
			l.tail = newNode
			l.head = l.tail
		} else {
			newNode.previous = l.tail
			newNode.previous.next = newNode
			l.tail = l.tail.next
		}
		l.CircularConnection()
		l.length++
	}
}

func (l *DoublyCircularLinked) IsEmpty() bool {
	return l.tail == nil
}

func (l *DoublyCircularLinked) IsOnlyOneNode() bool {
	return l.tail == l.head
}

func (l *DoublyCircularLinked) DeleteHead() error {
	if l.IsEmpty() {
		return fmt.Errorf("list is empty")
	} else {
		if l.IsOnlyOneNode() {
			l.head, l.tail = nil, nil
		} else {
			l.head = l.head.next
			l.head.previous = nil
			l.CircularConnection()
		}
	}
	l.length--
	return nil
}

func (l *DoublyCircularLinked) DeleteTail() error {
	if l.IsEmpty() {
		return fmt.Errorf("list is empty")
	} else {
		if l.IsOnlyOneNode() {
			l.tail, l.head = nil, nil
		} else {
			l.tail = l.tail.previous
			l.tail.next = nil
			l.CircularConnection()
		}
	}
	l.length--
	return nil
}

func (l *DoublyCircularLinked) Show() {
	if l.IsEmpty() {
		fmt.Errorf("list is empty")
		return
	}
	current := l.head
	for index := 0; index < l.length; index++ {
		fmt.Print(current.value, " -> ")
		current = current.next
	}
}

func (l *DoublyCircularLinked) ShowReverse() {
	if l.IsEmpty() {
		fmt.Errorf("list is empty")
		return
	}
	current := l.tail
	for index := 0; index < l.length; index++ {
		fmt.Print(current.value, " -> ")
		current = current.next
	}
}
