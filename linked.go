package main

import "fmt"

type linked struct {
	head   *node
	tail   *node
	length int
}

func (l *linked) insert(value int) {
	if err, newNode := newNode(value); err == nil {
		if l.tail == nil {
			l.head, l.tail = newNode, newNode
		} else {
			newNode.next = l.head
			newNode.next.previous = newNode
			l.head = l.head.previous
		}
		l.length++
	}
}

func (l *linked) deleteHead() error {
	if l.head == nil {
		return fmt.Errorf("list is empty")
	} else {
		l.head = l.head.next
		l.head.previous = nil
		return nil
	}
}

func (l *linked) append(value int) {
	if err, newNode := newNode(value); err == nil {
		if l.tail == nil {
			l.head, l.tail = newNode, newNode
		} else {
			newNode.previous = l.tail
			newNode.previous.next = newNode
			l.tail = l.tail.next
		}
		l.length++
	}
}

func (l *linked) deleteTail() error {
	if l.tail == nil {
		return fmt.Errorf("list is empty")
	} else {
		l.tail = l.tail.previous
		l.tail.next = nil
		return nil
	}
}

func (l *linked) show() {
	for current := l.head; current != nil; current = current.next {
		fmt.Print(current.value, " -> ")
	}
}

func (l *linked) showReverse() {
	for current := l.tail; current != nil; current = current.previous {
		fmt.Print(current.value, " -> ")
	}
}
